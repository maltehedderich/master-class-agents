package installer

import (
	"path/filepath"
	"testing"
)

func TestDefaultRegistryHasFourTools(t *testing.T) {
	r := DefaultRegistry()
	for _, name := range []string{"claude", "codex", "copilot", "gemini"} {
		if _, ok := r.Get(name); !ok {
			t.Errorf("DefaultRegistry missing %q", name)
		}
	}
}

func TestDefaultRegistryUnknown(t *testing.T) {
	r := DefaultRegistry()
	if _, ok := r.Get("unknown"); ok {
		t.Error("expected Get(unknown) to return false")
	}
}

func TestDefaultRegistryNames(t *testing.T) {
	r := DefaultRegistry()
	got := r.Names()
	want := map[string]bool{"claude": true, "codex": true, "copilot": true, "gemini": true}
	if len(got) != len(want) {
		t.Errorf("Names() len = %d, want %d", len(got), len(want))
	}
	for _, n := range got {
		if !want[n] {
			t.Errorf("unexpected installer name: %s", n)
		}
	}
}

func TestDefaultDirsUseHome(t *testing.T) {
	t.Setenv("HOME", "/tmp/fake-home")
	r := DefaultRegistry()

	cases := []struct {
		tool, kind, want string
	}{
		{"claude", "agent", filepath.FromSlash("/tmp/fake-home/.claude/agents")},
		{"claude", "skill", filepath.FromSlash("/tmp/fake-home/.claude/skills")},
		{"copilot", "agent", filepath.FromSlash("/tmp/fake-home/.copilot/agents")},
		{"copilot", "skill", filepath.FromSlash("/tmp/fake-home/.copilot/skills")},
		{"codex", "agent", filepath.FromSlash("/tmp/fake-home/.codex/skills")},
		{"codex", "skill", filepath.FromSlash("/tmp/fake-home/.codex/skills")},
		{"gemini", "agent", filepath.FromSlash("/tmp/fake-home/.gemini/skills")},
		{"gemini", "skill", filepath.FromSlash("/tmp/fake-home/.gemini/skills")},
	}
	for _, tc := range cases {
		t.Run(tc.tool+"-"+tc.kind, func(t *testing.T) {
			inst, ok := r.Get(tc.tool)
			if !ok {
				t.Fatalf("missing installer %q", tc.tool)
			}
			kind := KindAgent
			if tc.kind == "skill" {
				kind = KindSkill
			}
			if got := inst.DefaultDir(kind); got != tc.want {
				t.Errorf("DefaultDir(%s) = %q, want %q", tc.kind, got, tc.want)
			}
		})
	}
}
