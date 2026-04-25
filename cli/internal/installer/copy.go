package installer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

// homeJoin returns os.UserHomeDir() joined with elements. On error it falls
// back to "~" so DefaultDir always returns something printable.
func homeJoin(elem ...string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "~"
	}
	parts := append([]string{home}, elem...)
	return filepath.Join(parts...)
}

// copySkillDir mirrors src.Dir into destRoot/<src.Name>, applying conflict
// resolution per file via fsutil.WriteFile. Empty directories under the
// source skill are recreated at the destination.
func copySkillDir(src source.Skill, destRoot string, opts Options) (Result, error) {
	var res Result
	target := filepath.Join(destRoot, src.Name)

	err := filepath.WalkDir(src.Dir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(src.Dir, path)
		if err != nil {
			return err
		}
		out := filepath.Join(target, rel)

		if d.IsDir() {
			if !opts.DryRun {
				if err := os.MkdirAll(out, 0o755); err != nil {
					return fmt.Errorf("mkdir %s: %w", out, err)
				}
			}
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		one, err := writeOne(out, contents, opts)
		if err != nil {
			return err
		}
		res.Merge(one)
		return nil
	})
	if err != nil {
		return res, err
	}
	return res, nil
}
