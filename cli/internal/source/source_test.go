package source

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixtureRepo = "testdata/fixture-repo"

func TestResolveExplicitRepoFlag(t *testing.T) {
	abs, err := filepath.Abs(fixtureRepo)
	if err != nil {
		t.Fatal(err)
	}
	src, err := Resolve(abs)
	if err != nil {
		t.Fatalf("Resolve(%q) returned error: %v", abs, err)
	}
	if src.Root != abs {
		t.Errorf("Resolve.Root = %q, want %q", src.Root, abs)
	}
}

func TestResolveMissingRepo(t *testing.T) {
	tmp := t.TempDir()
	// chdir to an empty tree so the walk-up cannot find agents/ + skills/.
	t.Chdir(tmp)
	if _, err := Resolve(""); err == nil {
		t.Fatal("Resolve with no clone should return an error")
	}
}

func TestResolveExplicitFlagInvalid(t *testing.T) {
	tmp := t.TempDir() // empty: no agents/ or skills/
	if _, err := Resolve(tmp); err == nil {
		t.Fatalf("Resolve(%q) should fail because agents/ and skills/ are missing", tmp)
	}
}

func TestResolveWalksUpFromCwd(t *testing.T) {
	abs, err := filepath.Abs(fixtureRepo)
	if err != nil {
		t.Fatal(err)
	}
	// Start two levels deep inside the fixture repo.
	deep := filepath.Join(abs, "agents")
	t.Chdir(deep)

	src, err := Resolve("")
	if err != nil {
		t.Fatalf("Resolve walk-up failed: %v", err)
	}
	if src.Root != abs {
		t.Errorf("Resolve.Root = %q, want %q", src.Root, abs)
	}
}

func TestResolveEnvVar(t *testing.T) {
	abs, err := filepath.Abs(fixtureRepo)
	if err != nil {
		t.Fatal(err)
	}
	tmp := t.TempDir()
	t.Chdir(tmp)
	t.Setenv("MCAGENTS_REPO", abs)

	src, err := Resolve("")
	if err != nil {
		t.Fatalf("Resolve via env returned error: %v", err)
	}
	if src.Root != abs {
		t.Errorf("Resolve.Root = %q, want %q", src.Root, abs)
	}
}

func TestAgentsLoad(t *testing.T) {
	abs, _ := filepath.Abs(fixtureRepo)
	src := &Source{Root: abs}

	agents, err := src.Agents()
	if err != nil {
		t.Fatalf("Agents() error: %v", err)
	}
	if len(agents) != 1 {
		t.Fatalf("expected 1 agent in fixture, got %d", len(agents))
	}

	a := agents[0]
	if a.Name != "sample-agent" {
		t.Errorf("Agent.Name = %q, want sample-agent", a.Name)
	}
	if !strings.HasPrefix(a.Description, "Example agent used in mcagents tests") {
		t.Errorf("Agent.Description unexpected: %q", a.Description)
	}
	if a.Frontmatter["name"] != "Sample Agent" {
		t.Errorf("Frontmatter[name] = %q, want Sample Agent", a.Frontmatter["name"])
	}
	if !strings.Contains(a.Body, "You are a sample agent") {
		t.Errorf("Agent.Body missing expected content; got: %q", a.Body)
	}
	wantPath := filepath.Join(abs, "agents", "sample-agent.agent.md")
	if a.Path != wantPath {
		t.Errorf("Agent.Path = %q, want %q", a.Path, wantPath)
	}
}

func TestSkillsLoad(t *testing.T) {
	abs, _ := filepath.Abs(fixtureRepo)
	src := &Source{Root: abs}

	skills, err := src.Skills()
	if err != nil {
		t.Fatalf("Skills() error: %v", err)
	}
	if len(skills) != 1 {
		t.Fatalf("expected 1 skill, got %d", len(skills))
	}

	s := skills[0]
	if s.Name != "sample-skill" {
		t.Errorf("Skill.Name = %q, want sample-skill", s.Name)
	}
	if !strings.HasPrefix(s.Description, "Example skill used in mcagents tests") {
		t.Errorf("Skill.Description unexpected: %q", s.Description)
	}
	wantDir := filepath.Join(abs, "skills", "sample-skill")
	if s.Dir != wantDir {
		t.Errorf("Skill.Dir = %q, want %q", s.Dir, wantDir)
	}
}

func TestParseFrontmatterVariants(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  map[string]string
	}{
		{
			name: "quoted values",
			input: `---
name: "foo"
description: "Bar baz"
---
body
`,
			want: map[string]string{"name": "foo", "description": "Bar baz"},
		},
		{
			name: "unquoted values",
			input: `---
name: foo
description: bar baz
---
body
`,
			want: map[string]string{"name": "foo", "description": "bar baz"},
		},
		{
			name: "description with colons",
			input: `---
name: foo
description: "Use when: doing things; also when: doing other things."
---
body
`,
			want: map[string]string{
				"name":        "foo",
				"description": "Use when: doing things; also when: doing other things.",
			},
		},
		{
			name: "single quoted values",
			input: `---
name: 'foo'
description: 'bar'
---
body
`,
			want: map[string]string{"name": "foo", "description": "bar"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fm, _, err := parseFrontmatter([]byte(tc.input))
			if err != nil {
				t.Fatalf("parseFrontmatter error: %v", err)
			}
			for k, v := range tc.want {
				if fm[k] != v {
					t.Errorf("frontmatter[%q] = %q, want %q", k, fm[k], v)
				}
			}
		})
	}
}

func TestParseFrontmatterMissingFrontmatter(t *testing.T) {
	_, _, err := parseFrontmatter([]byte("just body, no frontmatter"))
	if err == nil {
		t.Fatal("expected error for missing frontmatter")
	}
}

func TestAgentMissingDescriptionErrors(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "agents"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "skills"), 0o755); err != nil {
		t.Fatal(err)
	}

	bad := filepath.Join(root, "agents", "no-desc.agent.md")
	contents := `---
name: "no-desc"
---
body
`
	if err := os.WriteFile(bad, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}

	src := &Source{Root: root}
	_, err := src.Agents()
	if err == nil {
		t.Fatal("expected error when agent description is missing")
	}
	if !strings.Contains(err.Error(), "no-desc") {
		t.Errorf("expected error to reference the file, got: %v", err)
	}
}

func TestAgentsEmptyDirReturnsEmpty(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "agents"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "skills"), 0o755); err != nil {
		t.Fatal(err)
	}
	src := &Source{Root: root}

	agents, err := src.Agents()
	if err != nil {
		t.Fatalf("Agents() empty-dir error: %v", err)
	}
	if len(agents) != 0 {
		t.Errorf("expected 0 agents, got %d", len(agents))
	}
	skills, err := src.Skills()
	if err != nil {
		t.Fatalf("Skills() empty-dir error: %v", err)
	}
	if len(skills) != 0 {
		t.Errorf("expected 0 skills, got %d", len(skills))
	}
}

func TestResolveErrorIsClear(t *testing.T) {
	tmp := t.TempDir()
	t.Chdir(tmp)

	_, err := Resolve("")
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, ErrSourceNotFound) {
		t.Errorf("expected ErrSourceNotFound, got %v", err)
	}
	if !strings.Contains(err.Error(), "agents/") || !strings.Contains(err.Error(), "skills/") {
		t.Errorf("expected error message to mention agents/ and skills/, got: %v", err)
	}
}
