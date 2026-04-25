package installer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestClaudeRewritesQuotedName(t *testing.T) {
	src := writeFixtureAgent(t, "backend-engineer", `---
name: "Backend Engineer"
description: "Sample description"
tools: [read]
---

Body
`)
	dest := t.TempDir()
	inst := &Claude{}
	if _, err := inst.InstallAgent(src, dest, Options{}); err != nil {
		t.Fatal(err)
	}

	got, _ := os.ReadFile(filepath.Join(dest, "backend-engineer.agent.md"))
	want := `name: "master-class-agents:backend-engineer"`
	if !strings.Contains(string(got), want) {
		t.Errorf("output missing %q; got:\n%s", want, got)
	}
	if strings.Contains(string(got), `name: "Backend Engineer"`) {
		t.Errorf("original name was not rewritten; got:\n%s", got)
	}
}

func TestClaudeRewritesUnquotedName(t *testing.T) {
	src := writeFixtureAgent(t, "foo", `---
name: foo-bar
description: "Sample"
---

Body
`)
	dest := t.TempDir()
	inst := &Claude{}
	if _, err := inst.InstallAgent(src, dest, Options{}); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(filepath.Join(dest, "foo.agent.md"))
	if !strings.Contains(string(got), `name: "master-class-agents:foo"`) {
		t.Errorf("output missing rewritten name; got:\n%s", got)
	}
}

func TestClaudeInsertsNameWhenMissing(t *testing.T) {
	src := writeFixtureAgent(t, "thing", `---
description: "Sample"
---

Body
`)
	dest := t.TempDir()
	inst := &Claude{}
	if _, err := inst.InstallAgent(src, dest, Options{}); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(filepath.Join(dest, "thing.agent.md"))
	if !strings.Contains(string(got), `name: "master-class-agents:thing"`) {
		t.Errorf("output missing inserted name; got:\n%s", got)
	}
	if !strings.Contains(string(got), `description: "Sample"`) {
		t.Errorf("description was lost; got:\n%s", got)
	}
}

func TestClaudeMultiLineFrontmatterPreserved(t *testing.T) {
	src := writeFixtureAgent(t, "multi", `---
name: "Multi"
description: "Sample"
tools:
  - read
  - edit
---

Body
`)
	dest := t.TempDir()
	inst := &Claude{}
	if _, err := inst.InstallAgent(src, dest, Options{}); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(filepath.Join(dest, "multi.agent.md"))
	out := string(got)
	if !strings.Contains(out, "  - read") || !strings.Contains(out, "  - edit") {
		t.Errorf("multi-line tools list lost; got:\n%s", out)
	}
	if !strings.Contains(out, `name: "master-class-agents:multi"`) {
		t.Errorf("missing rewritten name; got:\n%s", out)
	}
}

func TestClaudeInstallSkillCopiesVerbatim(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte("hi"), 0o644); err != nil {
		t.Fatal(err)
	}
	dest := t.TempDir()
	inst := &Claude{}
	skill := skillAt(dir, "demo")
	if _, err := inst.InstallSkill(skill, dest, Options{}); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(dest, "demo", "SKILL.md"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "hi" {
		t.Errorf("contents = %q, want hi", got)
	}
}

func TestClaudeName(t *testing.T) {
	if (&Claude{}).Name() != "claude" {
		t.Errorf("Name() = %q, want claude", (&Claude{}).Name())
	}
}
