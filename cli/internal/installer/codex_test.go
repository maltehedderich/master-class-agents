package installer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

func TestCodexInstallAgentWritesSkillFolder(t *testing.T) {
	agent := source.Agent{
		Name:        "backend-engineer",
		Description: "Use when: doing backend things",
		Path:        "/tmp/ignored",
		Frontmatter: map[string]string{"name": "Backend", "description": "ignored"},
		Body:        "\nYou are a backend engineer.\n\n## Principles\n",
	}
	dest := t.TempDir()
	inst := &Codex{}
	res, err := inst.InstallAgent(agent, dest, Options{})
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(dest, "backend-engineer", "SKILL.md")
	if !contains(res.Written, want) {
		t.Errorf("Written = %v, missing %s", res.Written, want)
	}
	got, err := os.ReadFile(want)
	if err != nil {
		t.Fatal(err)
	}
	out := string(got)
	if !strings.HasPrefix(out, "---\n") {
		t.Errorf("expected leading frontmatter delimiter; got:\n%s", out)
	}
	if !strings.Contains(out, `name: "backend-engineer"`) {
		t.Errorf("missing name in frontmatter; got:\n%s", out)
	}
	if !strings.Contains(out, `description: "Use when: doing backend things"`) {
		t.Errorf("description not preserved verbatim; got:\n%s", out)
	}
	if !strings.Contains(out, "You are a backend engineer.") {
		t.Errorf("body not preserved; got:\n%s", out)
	}
	if strings.Contains(out, `name: "Backend"`) {
		t.Errorf("original name leaked through; got:\n%s", out)
	}
}

func TestCodexEscapesQuotesInDescription(t *testing.T) {
	agent := source.Agent{
		Name:        "thing",
		Description: `Use "quoted" carefully`,
		Body:        "body\n",
		Frontmatter: map[string]string{},
	}
	dest := t.TempDir()
	inst := &Codex{}
	if _, err := inst.InstallAgent(agent, dest, Options{}); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(filepath.Join(dest, "thing", "SKILL.md"))
	if !strings.Contains(string(got), `description: "Use \"quoted\" carefully"`) {
		t.Errorf("quotes not escaped; got:\n%s", got)
	}
}

func TestCodexRejectsBadSkillName(t *testing.T) {
	agent := source.Agent{
		Name:        "Has Spaces",
		Description: "x",
		Body:        "y",
	}
	dest := t.TempDir()
	inst := &Codex{}
	if _, err := inst.InstallAgent(agent, dest, Options{}); err == nil {
		t.Fatal("expected error on bad skill name")
	}
}

func TestCodexInstallSkillCopiesVerbatim(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte("verbatim"), 0o644); err != nil {
		t.Fatal(err)
	}
	dest := t.TempDir()
	inst := &Codex{}
	skill := skillAt(dir, "demo")
	if _, err := inst.InstallSkill(skill, dest, Options{}); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(filepath.Join(dest, "demo", "SKILL.md"))
	if string(got) != "verbatim" {
		t.Errorf("got %q, want verbatim", got)
	}
}

func TestGeminiBehavesLikeCodex(t *testing.T) {
	agent := source.Agent{
		Name:        "foo",
		Description: "desc",
		Body:        "body\n",
	}
	dest := t.TempDir()
	inst := &Gemini{}
	if _, err := inst.InstallAgent(agent, dest, Options{}); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(filepath.Join(dest, "foo", "SKILL.md"))
	out := string(got)
	if !strings.Contains(out, `name: "foo"`) {
		t.Errorf("missing name; got:\n%s", out)
	}
	if !strings.Contains(out, `description: "desc"`) {
		t.Errorf("missing description; got:\n%s", out)
	}
}

func TestGeminiName(t *testing.T) {
	if (&Gemini{}).Name() != "gemini" {
		t.Errorf("Name() = %q, want gemini", (&Gemini{}).Name())
	}
}

func TestCodexName(t *testing.T) {
	if (&Codex{}).Name() != "codex" {
		t.Errorf("Name() = %q, want codex", (&Codex{}).Name())
	}
}

// skillAt is a small helper used across installer tests.
func skillAt(dir, name string) source.Skill {
	return source.Skill{Name: name, Description: "x", Dir: dir}
}
