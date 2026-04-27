package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/maltehedderich/master-class-agents/cli/internal/installer"
	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

type installFlags struct {
	tool      string
	agentsCSV string
	skillsCSV string
	dest      string
	force     bool
	noClobber bool
	repo      string
	dryRun    bool
	verbose   bool
}

func newInstallCmd(stdin io.Reader, stdout, stderr io.Writer) *cobra.Command {
	var flags installFlags

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install agents and skills into a target tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return newUsageError("install does not take positional arguments")
			}
			if flags.force && flags.noClobber {
				return newUsageError("--force and --no-clobber are mutually exclusive")
			}
			return runInstall(flags, stdin, stdout, stderr)
		},
	}

	cmd.Flags().StringVar(&flags.tool, "tool", "", "target tool: claude|codex|copilot|gemini")
	cmd.Flags().StringVar(&flags.agentsCSV, "agents", "", "comma-separated agent names (\"all\" for every agent)")
	cmd.Flags().StringVar(&flags.skillsCSV, "skills", "", "comma-separated skill names (\"all\" for every skill)")
	cmd.Flags().StringVar(&flags.dest, "dest", "", "override the tool's default destination")
	cmd.Flags().BoolVar(&flags.force, "force", false, "overwrite without prompting")
	cmd.Flags().BoolVar(&flags.noClobber, "no-clobber", false, "skip files that already exist")
	cmd.Flags().StringVar(&flags.repo, "repo", "", "override source repo location")
	cmd.Flags().BoolVar(&flags.dryRun, "dry-run", false, "print the plan; write nothing")
	cmd.Flags().BoolVarP(&flags.verbose, "verbose", "v", false, "log every file action")

	return cmd
}

func runInstall(flags installFlags, stdin io.Reader, stdout, stderr io.Writer) error {
	src, err := source.Resolve(flags.repo)
	if err != nil {
		return err
	}

	allAgents, err := src.Agents()
	if err != nil {
		return err
	}
	allSkills, err := src.Skills()
	if err != nil {
		return err
	}

	interactive := flags.agentsCSV == "" && flags.skillsCSV == ""

	if interactive {
		if !isTTY(stdin) {
			return newUsageError("non-interactive run requires --agents or --skills (and --tool)")
		}
		// Picker handles tool + selection.
		picked, err := runPicker(flags.tool, allAgents, allSkills, stdin, stdout, stderr)
		if err != nil {
			return err
		}
		if picked == nil {
			return nil
		}
		flags.tool = picked.tool
		// Replace catalogs with user-picked subsets.
		allAgents = picked.agents
		allSkills = picked.skills
		// In picker mode, both selections are explicit.
		flags.agentsCSV = "explicit"
		flags.skillsCSV = "explicit"
	}

	if flags.tool == "" {
		return newUsageError("--tool is required (one of claude, codex, copilot, gemini)")
	}

	registry := installer.DefaultRegistry()
	inst, ok := registry.Get(flags.tool)
	if !ok {
		return newUsageError("unknown tool %q (expected one of: %s)", flags.tool, strings.Join(registry.Names(), ", "))
	}

	selectedAgents, err := selectByName(allAgents, flags.agentsCSV, "agent", interactive)
	if err != nil {
		return err
	}
	selectedSkills, err := selectSkillsByName(allSkills, flags.skillsCSV, "skill", interactive)
	if err != nil {
		return err
	}

	if len(selectedAgents) == 0 && len(selectedSkills) == 0 {
		if err := writeLine(stderr, "nothing selected, exiting"); err != nil {
			return err
		}
		return nil
	}

	opts := installer.Options{
		Force:     flags.force,
		NoClobber: flags.noClobber,
		DryRun:    flags.dryRun,
	}
	if !flags.force && !flags.noClobber && interactive {
		opts.OnConflict = newInteractiveConflictHandler(stdin, stdout)
	}

	agentDest := flags.dest
	if agentDest == "" {
		agentDest = inst.DefaultDir(installer.KindAgent)
	}
	skillDest := flags.dest
	if skillDest == "" {
		skillDest = inst.DefaultDir(installer.KindSkill)
	}

	if flags.dryRun {
		if err := writeLine(stdout, "dry-run: planned writes"); err != nil {
			return err
		}
	}

	var totalRes installer.Result
	var failures []string
	for _, a := range selectedAgents {
		res, err := inst.InstallAgent(a, agentDest, opts)
		totalRes.Merge(res)
		if err != nil {
			failures = append(failures, fmt.Sprintf("agent %s: %v", a.Name, err))
			if err := writef(stderr, "FAIL agent %s: %v\n", a.Name, err); err != nil {
				return err
			}
			continue
		}
		if err := logActions(stdout, flags.verbose, flags.dryRun, "agent", a.Name, res); err != nil {
			return err
		}
	}
	for _, s := range selectedSkills {
		res, err := inst.InstallSkill(s, skillDest, opts)
		totalRes.Merge(res)
		if err != nil {
			failures = append(failures, fmt.Sprintf("skill %s: %v", s.Name, err))
			if err := writef(stderr, "FAIL skill %s: %v\n", s.Name, err); err != nil {
				return err
			}
			continue
		}
		if err := logActions(stdout, flags.verbose, flags.dryRun, "skill", s.Name, res); err != nil {
			return err
		}
	}

	verb := "Installed"
	if flags.dryRun {
		verb = "Would install"
	}
	if err := writef(stdout, "%s %d agent(s) and %d skill(s) into %q\n",
		verb, len(selectedAgents), len(selectedSkills), inst.Name()); err != nil {
		return err
	}
	if err := writef(stdout, "  %d file(s) written, %d skipped\n",
		len(totalRes.Written), len(totalRes.Skipped)); err != nil {
		return err
	}

	if len(failures) > 0 {
		return &artifactFailureError{
			msg: fmt.Sprintf("%d artifact(s) failed", len(failures)),
		}
	}
	return nil
}

func selectByName(all []source.Agent, csv, kind string, interactive bool) ([]source.Agent, error) {
	if csv == "" {
		return nil, nil
	}
	if csv == "all" {
		return all, nil
	}
	if interactive {
		// In picker mode the catalog has already been narrowed.
		return all, nil
	}
	byName := make(map[string]source.Agent, len(all))
	for _, a := range all {
		byName[a.Name] = a
	}
	var out []source.Agent
	for _, name := range splitCSV(csv) {
		a, ok := byName[name]
		if !ok {
			return nil, newUsageError("unknown %s %q", kind, name)
		}
		out = append(out, a)
	}
	return out, nil
}

func selectSkillsByName(all []source.Skill, csv, kind string, interactive bool) ([]source.Skill, error) {
	if csv == "" {
		return nil, nil
	}
	if csv == "all" {
		return all, nil
	}
	if interactive {
		return all, nil
	}
	byName := make(map[string]source.Skill, len(all))
	for _, s := range all {
		byName[s.Name] = s
	}
	var out []source.Skill
	for _, name := range splitCSV(csv) {
		s, ok := byName[name]
		if !ok {
			return nil, newUsageError("unknown %s %q", kind, name)
		}
		out = append(out, s)
	}
	return out, nil
}

func splitCSV(csv string) []string {
	parts := strings.Split(csv, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func logActions(w io.Writer, verbose, dryRun bool, kind, name string, res installer.Result) error {
	if !verbose && len(res.Written) <= 1 && len(res.Skipped) == 0 {
		marker := "installed"
		if dryRun {
			marker = "would install"
		}
		return writef(w, "%s %s %s\n", marker, kind, name)
	}
	for _, p := range res.Written {
		marker := "wrote"
		if dryRun {
			marker = "would write"
		}
		if err := writef(w, "  %s %s\n", marker, p); err != nil {
			return err
		}
	}
	for _, p := range res.Skipped {
		if err := writef(w, "  skipped %s (already up to date)\n", p); err != nil {
			return err
		}
	}
	return nil
}

func isTTY(r io.Reader) bool {
	type fdHolder interface{ Fd() uintptr }
	f, ok := r.(fdHolder)
	if !ok {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}
