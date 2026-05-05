package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// fixtureRepo points at the source-package fixture clone so the CLI tests
// don't drift from the loader tests.
func fixtureRepo(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// internal/cli -> internal/source/testdata/fixture-repo
	repo := filepath.Join(wd, "..", "source", "testdata", "fixture-repo")
	abs, err := filepath.Abs(repo)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(abs, "agents")); err != nil {
		t.Fatalf("fixture repo missing: %v", err)
	}
	return abs
}

func runCLI(t *testing.T, args ...string) (int, string, string) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	code := Run(args, nil, &stdout, &stderr)
	return code, stdout.String(), stderr.String()
}

func TestInstallCopilotSingleAgent(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, errOut := runCLI(t,
		"install",
		"--tool", "copilot",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr:\n%s", code, errOut)
	}
	want := filepath.Join(dest, "sample-agent.agent.md")
	if _, err := os.Stat(want); err != nil {
		t.Fatalf("expected %s to exist: %v", want, err)
	}
}

func TestInstallAllAgents(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, errOut := runCLI(t,
		"install",
		"--tool", "copilot",
		"--agents", "all",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr:\n%s", code, errOut)
	}
	if _, err := os.Stat(filepath.Join(dest, "sample-agent.agent.md")); err != nil {
		t.Fatalf("missing installed agent: %v", err)
	}
}

func TestInstallClaudeRewritesName(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, _ := runCLI(t,
		"install",
		"--tool", "claude",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	got, err := os.ReadFile(filepath.Join(dest, "sample-agent.agent.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), `name: "master-class-agents:sample-agent"`) {
		t.Errorf("missing namespaced name; got:\n%s", got)
	}
}

func TestInstallCodexAgentEmitsNativeAgentFile(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, _ := runCLI(t,
		"install",
		"--tool", "codex",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	target := filepath.Join(dest, "sample-agent.toml")
	if _, err := os.Stat(target); err != nil {
		t.Fatalf("missing native Codex agent file: %v", err)
	}
}

func TestInstallSkill(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, _ := runCLI(t,
		"install",
		"--tool", "copilot",
		"--skills", "sample-skill",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if _, err := os.Stat(filepath.Join(dest, "sample-skill", "SKILL.md")); err != nil {
		t.Fatalf("missing skill: %v", err)
	}
}

func TestInstallNoClobberLeavesExistingDifferingFile(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()
	existing := filepath.Join(dest, "sample-agent.agent.md")
	if err := os.WriteFile(existing, []byte("old content"), 0o644); err != nil {
		t.Fatal(err)
	}

	code, _, _ := runCLI(t,
		"install",
		"--tool", "copilot",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
		"--no-clobber",
	)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	got, _ := os.ReadFile(existing)
	if string(got) != "old content" {
		t.Errorf("file changed despite --no-clobber: %q", got)
	}
}

func TestInstallForceOverwrites(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()
	existing := filepath.Join(dest, "sample-agent.agent.md")
	if err := os.WriteFile(existing, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	code, _, _ := runCLI(t,
		"install",
		"--tool", "copilot",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
		"--force",
	)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	got, _ := os.ReadFile(existing)
	if string(got) == "old" {
		t.Error("file was not overwritten")
	}
}

func TestInstallDryRunWritesNothing(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, out, _ := runCLI(t,
		"install",
		"--tool", "copilot",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
		"--dry-run",
	)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if _, err := os.Stat(filepath.Join(dest, "sample-agent.agent.md")); !os.IsNotExist(err) {
		t.Errorf("dry-run wrote file; stat err = %v", err)
	}
	if !strings.Contains(out, "dry-run") && !strings.Contains(out, "would") {
		t.Errorf("dry-run output should announce itself; got:\n%s", out)
	}
}

func TestInstallUnknownToolExits2(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, errOut := runCLI(t,
		"install",
		"--tool", "bogus",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 2 {
		t.Errorf("exit = %d, want 2", code)
	}
	if !strings.Contains(errOut, "bogus") {
		t.Errorf("stderr should mention bogus tool: %q", errOut)
	}
}

func TestInstallUnknownAgentExits2(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, errOut := runCLI(t,
		"install",
		"--tool", "copilot",
		"--agents", "does-not-exist",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 2 {
		t.Errorf("exit = %d, want 2", code)
	}
	if !strings.Contains(errOut, "does-not-exist") {
		t.Errorf("stderr should name the unknown agent: %q", errOut)
	}
}

func TestInstallNoTTYNoToolExits2(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, errOut := runCLI(t,
		"install",
		"--repo", repo,
		"--dest", dest,
	)
	if code != 2 {
		t.Errorf("exit = %d, want 2", code)
	}
	if !strings.Contains(strings.ToLower(errOut), "tool") {
		t.Errorf("stderr should ask for --tool: %q", errOut)
	}
}

func TestInstallForceAndNoClobberConflict(t *testing.T) {
	repo := fixtureRepo(t)
	dest := t.TempDir()

	code, _, errOut := runCLI(t,
		"install",
		"--tool", "copilot",
		"--agents", "sample-agent",
		"--repo", repo,
		"--dest", dest,
		"--force",
		"--no-clobber",
	)
	if code != 2 {
		t.Errorf("exit = %d, want 2", code)
	}
	if !strings.Contains(strings.ToLower(errOut), "force") {
		t.Errorf("stderr should explain mutex; got: %q", errOut)
	}
}
