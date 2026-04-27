// Package fsutil writes files atomically with conflict-resolution semantics
// shared by every installer.
package fsutil

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ErrAborted is returned when WriteFile is told to abort by an OnConflict
// callback or by an installer that propagates user cancellation.
var ErrAborted = errors.New("install aborted")

// Action describes what WriteFile did (or what an OnConflict callback wants
// it to do).
type Action int

const (
	// ActionOverwrite means the destination was (or should be) written.
	ActionOverwrite Action = iota
	// ActionSkip means the destination was left unchanged.
	ActionSkip
	// ActionAbort means the caller should stop the entire install run.
	ActionAbort
)

// String returns a stable human-readable label for an Action.
func (a Action) String() string {
	switch a {
	case ActionOverwrite:
		return "overwrite"
	case ActionSkip:
		return "skip"
	case ActionAbort:
		return "abort"
	default:
		return fmt.Sprintf("Action(%d)", int(a))
	}
}

// Options controls conflict resolution and dry-run behaviour for WriteFile
// (and, by extension, every Installer).
type Options struct {
	// Force overwrites differing files without prompting.
	Force bool
	// NoClobber skips differing files without writing. NoClobber wins over
	// Force when both are set.
	NoClobber bool
	// DryRun computes the Action and reports what would happen without
	// touching the filesystem.
	DryRun bool
	// OnConflict, if non-nil, is consulted whenever a destination exists with
	// content differing from the incoming bytes (and neither Force nor
	// NoClobber is set). The closure typically prompts the user.
	OnConflict func(path string, existing, incoming []byte) Action
}

// WriteFile writes content to dest with atomic semantics:
//
//   - If dest does not exist, it is created.
//   - If dest exists with identical content, it is left alone (ActionSkip).
//   - If dest exists with different content, the resolution depends on opts:
//     NoClobber → skip; Force → overwrite; OnConflict → ask; default → write.
//   - DryRun computes the Action without writing.
//
// All writes go through a sibling .tmp file followed by os.Rename so an
// interrupted run never leaves a half-written destination.
func WriteFile(dest string, content []byte, opts Options) (Action, error) {
	existing, exists, err := readIfExists(dest)
	if err != nil {
		return 0, err
	}

	if exists && bytes.Equal(existing, content) {
		return ActionSkip, nil
	}

	if exists {
		switch {
		case opts.NoClobber:
			return ActionSkip, nil
		case opts.Force:
			// fall through to write
		case opts.OnConflict != nil:
			act := opts.OnConflict(dest, existing, content)
			switch act {
			case ActionSkip:
				return ActionSkip, nil
			case ActionAbort:
				return ActionAbort, ErrAborted
			case ActionOverwrite:
				// fall through
			default:
				return 0, fmt.Errorf("OnConflict returned unknown action: %v", act)
			}
		}
	}

	if opts.DryRun {
		return ActionOverwrite, nil
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return 0, fmt.Errorf("mkdir parent: %w", err)
	}

	if err := atomicWrite(dest, content); err != nil {
		return 0, err
	}
	return ActionOverwrite, nil
}

func readIfExists(path string) ([]byte, bool, error) {
	data, err := os.ReadFile(path)
	if err == nil {
		return data, true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return nil, false, nil
	}
	return nil, false, fmt.Errorf("read %s: %w", path, err)
}

func atomicWrite(dest string, content []byte) error {
	tmp := dest + ".tmp"
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("open temp %s: %w", tmp, err)
	}
	if _, err := f.Write(content); err != nil {
		return errors.Join(
			fmt.Errorf("write temp %s: %w", tmp, err),
			closeTemp(f, tmp),
			removeTemp(tmp),
		)
	}
	if err := f.Sync(); err != nil {
		return errors.Join(
			fmt.Errorf("fsync temp %s: %w", tmp, err),
			closeTemp(f, tmp),
			removeTemp(tmp),
		)
	}
	if err := f.Close(); err != nil {
		return errors.Join(
			fmt.Errorf("close temp %s: %w", tmp, err),
			removeTemp(tmp),
		)
	}
	if err := os.Rename(tmp, dest); err != nil {
		return errors.Join(
			fmt.Errorf("rename %s -> %s: %w", tmp, dest, err),
			removeTemp(tmp),
		)
	}
	return nil
}

func closeTemp(f *os.File, path string) error {
	if err := f.Close(); err != nil {
		return fmt.Errorf("close temp %s: %w", path, err)
	}
	return nil
}

func removeTemp(path string) error {
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove temp %s: %w", path, err)
	}
	return nil
}
