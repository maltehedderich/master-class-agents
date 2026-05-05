package cli

import (
	"encoding/json"
	"io"
	"path/filepath"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/maltehedderich/master-class-agents/cli/internal/installer"
	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

type listFlags struct {
	tool   string
	repo   string
	asJSON bool
}

func newListCmd(stdout, _ io.Writer) *cobra.Command {
	var flags listFlags
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available agents and skills",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(flags, stdout)
		},
	}
	cmd.Flags().StringVar(&flags.tool, "tool", "", "show destination column for that tool")
	cmd.Flags().StringVar(&flags.repo, "repo", "", "override source repo location")
	cmd.Flags().BoolVar(&flags.asJSON, "json", false, "emit JSON instead of a table")
	return cmd
}

type listEntry struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Dest        string `json:"dest,omitempty"`
}

func runList(flags listFlags, stdout io.Writer) error {
	src, err := source.Resolve(flags.repo)
	if err != nil {
		return err
	}
	agents, err := src.Agents()
	if err != nil {
		return err
	}
	skills, err := src.Skills()
	if err != nil {
		return err
	}

	var inst installer.Installer
	if flags.tool != "" {
		registry := installer.DefaultRegistry()
		got, ok := registry.Get(flags.tool)
		if !ok {
			return newUsageError("unknown tool %q", flags.tool)
		}
		inst = got
	}

	entries := make([]listEntry, 0, len(agents)+len(skills))
	for _, a := range agents {
		e := listEntry{Name: a.Name, Type: "agent", Description: a.Description}
		if inst != nil {
			e.Dest = agentDest(inst, a)
		}
		entries = append(entries, e)
	}
	for _, s := range skills {
		e := listEntry{Name: s.Name, Type: "skill", Description: s.Description}
		if inst != nil {
			e.Dest = filepath.Join(inst.DefaultDir(installer.KindSkill), s.Name)
		}
		entries = append(entries, e)
	}

	if flags.asJSON {
		enc := json.NewEncoder(stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(entries)
	}

	tw := tabwriter.NewWriter(stdout, 0, 0, 2, ' ', 0)
	if inst != nil {
		if err := writeLine(tw, "NAME\tTYPE\tDESCRIPTION\tDEST"); err != nil {
			return err
		}
	} else {
		if err := writeLine(tw, "NAME\tTYPE\tDESCRIPTION"); err != nil {
			return err
		}
	}
	for _, e := range entries {
		desc := e.Description
		if len(desc) > 80 {
			desc = desc[:77] + "..."
		}
		if inst != nil {
			if err := writef(tw, "%s\t%s\t%s\t%s\n", e.Name, e.Type, desc, e.Dest); err != nil {
				return err
			}
		} else {
			if err := writef(tw, "%s\t%s\t%s\n", e.Name, e.Type, desc); err != nil {
				return err
			}
		}
	}
	return tw.Flush()
}

func agentDest(inst installer.Installer, a source.Agent) string {
	root := inst.DefaultDir(installer.KindAgent)
	switch inst.Name() {
	case "codex":
		return filepath.Join(root, a.Name+".toml")
	case "gemini":
		return filepath.Join(root, a.Name, "SKILL.md")
	default:
		return filepath.Join(root, a.Name+".agent.md")
	}
}
