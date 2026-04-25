// Package installer turns parsed agents and skills into on-disk artifacts
// for a target tool (Claude, Codex, Copilot, Gemini).
//
// Each tool has its own implementation in claude.go / codex.go / copilot.go /
// gemini.go. They share the WriteFile primitive in internal/fsutil so that
// conflict resolution, dry-run, and atomic writes behave identically across
// tools.
package installer

import (
	"github.com/maltehedderich/master-class-agents/cli/internal/fsutil"
	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

// Kind distinguishes installable artifact types.
type Kind int

const (
	// KindAgent is an agent file from agents/.
	KindAgent Kind = iota
	// KindSkill is a skill folder from skills/.
	KindSkill
)

// Options re-exports fsutil.Options so callers don't need both packages.
type Options = fsutil.Options

// Action re-exports fsutil.Action.
type Action = fsutil.Action

const (
	// ActionOverwrite re-exports fsutil.ActionOverwrite.
	ActionOverwrite = fsutil.ActionOverwrite
	// ActionSkip re-exports fsutil.ActionSkip.
	ActionSkip = fsutil.ActionSkip
	// ActionAbort re-exports fsutil.ActionAbort.
	ActionAbort = fsutil.ActionAbort
)

// Result is the per-call summary of files an installer touched.
type Result struct {
	Written []string
	Skipped []string
}

// Merge appends other into r in place.
func (r *Result) Merge(other Result) {
	r.Written = append(r.Written, other.Written...)
	r.Skipped = append(r.Skipped, other.Skipped...)
}

// Installer writes agents and skills for one target tool.
type Installer interface {
	// Name returns the canonical tool name ("claude", "codex", ...).
	Name() string
	// DefaultDir returns the per-kind default destination directory under
	// $HOME for this tool.
	DefaultDir(kind Kind) string
	// InstallAgent writes an agent into destRoot using opts. The destination
	// path is implementation-specific.
	InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error)
	// InstallSkill writes a skill into destRoot using opts.
	InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error)
}
