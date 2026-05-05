package installer

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

// Codex installs agents as native Codex custom-agent TOML files. Skills
// install verbatim.
type Codex struct{}

// Name returns "codex".
func (*Codex) Name() string { return "codex" }

// DefaultDir returns ~/.codex/agents or ~/.codex/skills.
func (*Codex) DefaultDir(kind Kind) string {
	switch kind {
	case KindAgent:
		return homeJoin(".codex", "agents")
	default:
		return homeJoin(".codex", "skills")
	}
}

// InstallAgent emits destRoot/<a.Name>.toml using Codex's custom-agent schema.
func (*Codex) InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error) {
	return writeCodexAgent(a, destRoot, opts)
}

// InstallSkill recursively copies the source skill folder under destRoot.
func (*Codex) InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error) {
	return copySkillDir(s, destRoot, opts)
}

// Gemini installs agents as Gemini skill folders and writes under ~/.gemini.
type Gemini struct{}

// Name returns "gemini".
func (*Gemini) Name() string { return "gemini" }

// DefaultDir returns ~/.gemini/skills for both kinds.
func (*Gemini) DefaultDir(kind Kind) string {
	return homeJoin(".gemini", "skills")
}

// InstallAgent emits destRoot/<a.Name>/SKILL.md with a fresh frontmatter
// containing the agent name and description, followed by the agent body.
func (*Gemini) InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error) {
	return writeAgentAsSkill(a, destRoot, opts)
}

// InstallSkill recursively copies the source skill folder under destRoot.
func (*Gemini) InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error) {
	return copySkillDir(s, destRoot, opts)
}

var artifactNameRE = regexp.MustCompile(`^[a-z0-9-]+$`)

func writeCodexAgent(a source.Agent, destRoot string, opts Options) (Result, error) {
	if !artifactNameRE.MatchString(a.Name) {
		return Result{}, fmt.Errorf("invalid agent name %q (must match [a-z0-9-]+)", a.Name)
	}
	if a.Description == "" {
		return Result{}, fmt.Errorf("agent %q has no description", a.Name)
	}

	body := strings.TrimLeft(a.Body, "\n")

	var b strings.Builder
	fmt.Fprintf(&b, "name = %s\n", tomlBasicString(a.Name))
	fmt.Fprintf(&b, "description = %s\n", tomlBasicString(a.Description))
	fmt.Fprintf(&b, "developer_instructions = %s\n", tomlInstructionString(body))

	dest := filepath.Join(destRoot, a.Name+".toml")
	return writeOne(dest, []byte(b.String()), opts)
}

func writeAgentAsSkill(a source.Agent, destRoot string, opts Options) (Result, error) {
	if !artifactNameRE.MatchString(a.Name) {
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

func tomlInstructionString(s string) string {
	if !strings.Contains(s, "'''") {
		return "'''\n" + s + "'''"
	}
	return tomlBasicString(s)
}

func tomlBasicString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r == utf8.RuneError && size == 1 {
			b.WriteString(`\uFFFD`)
			s = s[size:]
			continue
		}
		switch r {
		case '\b':
			b.WriteString(`\b`)
		case '\t':
			b.WriteString(`\t`)
		case '\n':
			b.WriteString(`\n`)
		case '\f':
			b.WriteString(`\f`)
		case '\r':
			b.WriteString(`\r`)
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		default:
			switch {
			case r < 0x20 || r == 0x7f:
				fmt.Fprintf(&b, `\u%04X`, r)
			case r <= 0xffff:
				b.WriteRune(r)
			default:
				b.WriteRune(r)
			}
		}
		s = s[size:]
	}
	b.WriteByte('"')
	return b.String()
}
