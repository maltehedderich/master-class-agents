// Package source resolves the master-class-agents repo on disk and loads
// agent and skill catalogs from it.
package source

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ErrSourceNotFound is returned by Resolve when no master-class-agents repo
// can be located via the flag, env var, or walking up from cwd.
var ErrSourceNotFound = errors.New("master-class-agents source not found")

// EnvVar is the environment variable consulted by Resolve as the second
// fallback after the explicit flag.
const EnvVar = "MCAGENTS_REPO"

// Source is a located clone of the master-class-agents repo.
type Source struct {
	Root string
}

// Resolve picks the source repo to read agents/ and skills/ from.
//
// Resolution order:
//  1. The explicit repoFlag argument, if non-empty.
//  2. The MCAGENTS_REPO environment variable.
//  3. Walking up from os.Getwd until a directory containing both
//     agents/ and skills/ is found.
//
// On failure the returned error wraps ErrSourceNotFound and includes a
// human-actionable message.
func Resolve(repoFlag string) (*Source, error) {
	if repoFlag != "" {
		return validate(repoFlag)
	}
	if env := os.Getenv(EnvVar); env != "" {
		return validate(env)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("getwd: %w", err)
	}

	dir := cwd
	for {
		if hasRepoLayout(dir) {
			abs, err := filepath.Abs(dir)
			if err != nil {
				return nil, err
			}
			return &Source{Root: abs}, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return nil, fmt.Errorf(
		"%w: cannot find master-class-agents repo (looked for agents/ and skills/ from %s up to /). "+
			"Run from a clone, or pass --repo PATH or set %s",
		ErrSourceNotFound, cwd, EnvVar,
	)
}

func validate(path string) (*Source, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	if !hasRepoLayout(abs) {
		return nil, fmt.Errorf(
			"%w: %s does not contain agents/ and skills/ subdirectories",
			ErrSourceNotFound, abs,
		)
	}
	return &Source{Root: abs}, nil
}

func hasRepoLayout(dir string) bool {
	return isDir(filepath.Join(dir, "agents")) && isDir(filepath.Join(dir, "skills"))
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
