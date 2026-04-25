package cli

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestListPrintsAllArtifacts(t *testing.T) {
	repo := fixtureRepo(t)

	code, out, _ := runCLI(t, "list", "--repo", repo)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if !strings.Contains(out, "sample-agent") {
		t.Errorf("missing sample-agent in output:\n%s", out)
	}
	if !strings.Contains(out, "sample-skill") {
		t.Errorf("missing sample-skill in output:\n%s", out)
	}
}

func TestListWithToolShowsDest(t *testing.T) {
	repo := fixtureRepo(t)

	code, out, _ := runCLI(t, "list", "--tool", "claude", "--repo", repo)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if !strings.Contains(out, ".claude") {
		t.Errorf("expected destination column to mention .claude; got:\n%s", out)
	}
}

func TestListJSON(t *testing.T) {
	repo := fixtureRepo(t)

	code, out, _ := runCLI(t, "list", "--repo", repo, "--json")
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	var entries []map[string]any
	if err := json.Unmarshal([]byte(out), &entries); err != nil {
		t.Fatalf("not valid JSON: %v\n%s", err, out)
	}
	if len(entries) < 2 {
		t.Errorf("expected at least 2 entries, got %d", len(entries))
	}
	for _, e := range entries {
		if e["name"] == nil {
			t.Errorf("entry missing name: %v", e)
		}
		if e["type"] == nil {
			t.Errorf("entry missing type: %v", e)
		}
	}
}
