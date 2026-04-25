package cli

import (
	"fmt"
	"io"

	"github.com/charmbracelet/huh"

	"github.com/maltehedderich/master-class-agents/cli/internal/installer"
	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

type pickerResult struct {
	tool   string
	agents []source.Agent
	skills []source.Skill
}

// runPicker presents the interactive form. The stdin/stdout arguments are
// reserved for huh.Form.WithInput/WithOutput; huh currently writes to the
// real terminal so test runs (TTY=false) skip this path entirely.
func runPicker(initialTool string, agents []source.Agent, skills []source.Skill, _ io.Reader, _, _ io.Writer) (*pickerResult, error) {
	tool := initialTool
	var agentNames, skillNames []string

	registry := installer.DefaultRegistry()
	toolNames := registry.Names()

	groups := []*huh.Group{}
	if tool == "" {
		toolOpts := make([]huh.Option[string], 0, len(toolNames))
		for _, n := range toolNames {
			toolOpts = append(toolOpts, huh.NewOption(n, n))
		}
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[string]().
				Title("Target tool").
				Options(toolOpts...).
				Value(&tool),
		))
	}

	if len(agents) > 0 {
		agentOpts := make([]huh.Option[string], 0, len(agents))
		for _, a := range agents {
			agentOpts = append(agentOpts, huh.NewOption(fmt.Sprintf("%s — %s", a.Name, truncate(a.Description, 60)), a.Name))
		}
		groups = append(groups, huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Agents to install").
				Options(agentOpts...).
				Value(&agentNames),
		))
	}

	if len(skills) > 0 {
		skillOpts := make([]huh.Option[string], 0, len(skills))
		for _, s := range skills {
			skillOpts = append(skillOpts, huh.NewOption(fmt.Sprintf("%s — %s", s.Name, truncate(s.Description, 60)), s.Name))
		}
		groups = append(groups, huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Skills to install").
				Options(skillOpts...).
				Value(&skillNames),
		))
	}

	form := huh.NewForm(groups...)
	if err := form.Run(); err != nil {
		return nil, err
	}

	if _, ok := registry.Get(tool); !ok {
		return nil, newUsageError("unknown tool %q", tool)
	}

	selectedAgents := filterAgents(agents, agentNames)
	selectedSkills := filterSkills(skills, skillNames)

	return &pickerResult{
		tool:   tool,
		agents: selectedAgents,
		skills: selectedSkills,
	}, nil
}

func filterAgents(all []source.Agent, names []string) []source.Agent {
	if len(names) == 0 {
		return nil
	}
	wanted := make(map[string]bool, len(names))
	for _, n := range names {
		wanted[n] = true
	}
	var out []source.Agent
	for _, a := range all {
		if wanted[a.Name] {
			out = append(out, a)
		}
	}
	return out
}

func filterSkills(all []source.Skill, names []string) []source.Skill {
	if len(names) == 0 {
		return nil
	}
	wanted := make(map[string]bool, len(names))
	for _, n := range names {
		wanted[n] = true
	}
	var out []source.Skill
	for _, s := range all {
		if wanted[s.Name] {
			out = append(out, s)
		}
	}
	return out
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-1] + "…"
}
