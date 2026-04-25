package installer

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

// Codex installs every agent as a Codex skill folder. Skills install
// verbatim.
type Codex struct{}

// Name returns "codex".
func (*Codex) Name() string { return "codex" }

// DefaultDir returns ~/.codex/skills for both kinds.
func (*Codex) DefaultDir(kind Kind) string {
	return homeJoin(".codex", "skills")
}

// InstallAgent emits destRoot/<a.Name>/SKILL.md with a fresh frontmatter
// containing the agent name and description, followed by the agent body.
func (*Codex) InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error) {
	return writeAgentAsSkill(a, destRoot, opts)
}

// InstallSkill recursively copies the source skill folder under destRoot.
func (*Codex) InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error) {
	return copySkillDir(s, destRoot, opts)
}

// Gemini matches Codex behaviour for both agents and skills, but writes
// under ~/.gemini.
type Gemini struct{}

// Name returns "gemini".
func (*Gemini) Name() string { return "gemini" }

// DefaultDir returns ~/.gemini/skills for both kinds.
func (*Gemini) DefaultDir(kind Kind) string {
	return homeJoin(".gemini", "skills")
}

// InstallAgent behaves like Codex.InstallAgent.
func (*Gemini) InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error) {
	return writeAgentAsSkill(a, destRoot, opts)
}

// InstallSkill recursively copies the source skill folder under destRoot.
func (*Gemini) InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error) {
	return copySkillDir(s, destRoot, opts)
}

var skillNameRE = regexp.MustCompile(`^[a-z0-9-]+$`)

func writeAgentAsSkill(a source.Agent, destRoot string, opts Options) (Result, error) {
	if !skillNameRE.MatchString(a.Name) {
		return Result{}, fmt.Errorf("invalid skill name %q (must match [a-z0-9-]+)", a.Name)
	}
	if a.Description == "" {
		return Result{}, fmt.Errorf("agent %q has no description", a.Name)
	}

	body := strings.TrimLeft(a.Body, "\n")

	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "name: %q\n", a.Name)
	fmt.Fprintf(&b, "description: %q\n", a.Description)
	b.WriteString("---\n")
	b.WriteString(body)

	dest := filepath.Join(destRoot, a.Name, "SKILL.md")
	return writeOne(dest, []byte(b.String()), opts)
}
