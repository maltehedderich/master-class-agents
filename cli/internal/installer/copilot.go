package installer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/maltehedderich/master-class-agents/cli/internal/fsutil"
	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

// Copilot installs agents and skills as-is for GitHub Copilot.
type Copilot struct{}

// Name returns "copilot".
func (*Copilot) Name() string { return "copilot" }

// DefaultDir returns ~/.copilot/agents or ~/.copilot/skills.
func (*Copilot) DefaultDir(kind Kind) string {
	switch kind {
	case KindAgent:
		return homeJoin(".copilot", "agents")
	default:
		return homeJoin(".copilot", "skills")
	}
}

// InstallAgent copies the agent file byte-for-byte into destRoot/<name>.agent.md.
func (*Copilot) InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error) {
	contents, err := os.ReadFile(a.Path)
	if err != nil {
		return Result{}, fmt.Errorf("read agent %s: %w", a.Path, err)
	}
	dest := filepath.Join(destRoot, a.Name+".agent.md")
	return writeOne(dest, contents, opts)
}

// InstallSkill recursively copies the source skill folder under destRoot.
func (*Copilot) InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error) {
	return copySkillDir(s, destRoot, opts)
}

// writeOne is the single-file write helper used by every installer.
func writeOne(dest string, contents []byte, opts Options) (Result, error) {
	action, err := fsutil.WriteFile(dest, contents, opts)
	if err != nil {
		return Result{}, err
	}
	switch action {
	case fsutil.ActionOverwrite:
		return Result{Written: []string{dest}}, nil
	case fsutil.ActionSkip:
		return Result{Skipped: []string{dest}}, nil
	default:
		return Result{}, nil
	}
}
