package installer

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

func TestCopilotInstallAgentCopiesAsIs(t *testing.T) {
	src := writeFixtureAgent(t, "sample", `---
name: "Sample Agent"
description: "Sample description"
---

Body
`)
	dest := t.TempDir()

	inst := &Copilot{}
	res, err := inst.InstallAgent(src, dest, Options{})
	if err != nil {
		t.Fatalf("InstallAgent: %v", err)
	}
	want := filepath.Join(dest, "sample.agent.md")
	if !contains(res.Written, want) {
		t.Errorf("Written = %v, missing %s", res.Written, want)
	}

	got, err := os.ReadFile(want)
	if err != nil {
		t.Fatal(err)
	}
	original, _ := os.ReadFile(src.Path)
	if !bytes.Equal(got, original) {
		t.Errorf("copilot install changed bytes; got %q, want %q", got, original)
	}
}

func TestCopilotInstallSkillCopiesDirVerbatim(t *testing.T) {
	skillDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(skillDir, "ref"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "ref", "more.md"), []byte("ref"), 0o644); err != nil {
		t.Fatal(err)
	}

	dest := t.TempDir()
	inst := &Copilot{}
	skill := source.Skill{Name: "demo", Description: "Demo skill", Dir: skillDir}
	res, err := inst.InstallSkill(skill, dest, Options{})
	if err != nil {
		t.Fatalf("InstallSkill: %v", err)
	}
	if len(res.Written) < 2 {
		t.Errorf("Written should include nested file; got %v", res.Written)
	}

	got, _ := os.ReadFile(filepath.Join(dest, "demo", "SKILL.md"))
	if string(got) != "hello" {
		t.Errorf("SKILL.md contents = %q, want hello", got)
	}
	got, _ = os.ReadFile(filepath.Join(dest, "demo", "ref", "more.md"))
	if string(got) != "ref" {
		t.Errorf("nested file contents = %q, want ref", got)
	}
}

func TestCopilotName(t *testing.T) {
	if (&Copilot{}).Name() != "copilot" {
		t.Errorf("Name() = %q, want copilot", (&Copilot{}).Name())
	}
}

// writeFixtureAgent writes a temporary agent file and returns a source.Agent
// pointing at it.
func writeFixtureAgent(t *testing.T, name, contents string) source.Agent {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, name+".agent.md")
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}
	return source.Agent{
		Name:        name,
		Description: "Sample description",
		Path:        path,
		Frontmatter: map[string]string{
			"name":        name,
			"description": "Sample description",
		},
		Body: "\nBody\n",
	}
}

func contains(haystack []string, needle string) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}
	return false
}
