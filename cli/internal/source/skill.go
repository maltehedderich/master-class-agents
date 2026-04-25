package source

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// Skill is a parsed skills/<name>/SKILL.md artifact.
type Skill struct {
	Name        string
	Description string
	Dir         string
}

// Skills loads every skills/<name>/SKILL.md under s.Root.
//
// Each skill folder must contain a SKILL.md with a non-empty description.
// Folders without a SKILL.md are skipped silently to allow drafts.
func (s *Source) Skills() ([]Skill, error) {
	dir := filepath.Join(s.Root, "skills")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read skills dir %s: %w", dir, err)
	}

	var skills []Skill
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		skillDir := filepath.Join(dir, entry.Name())
		skillFile := filepath.Join(skillDir, "SKILL.md")
		info, err := os.Stat(skillFile)
		if err != nil || info.IsDir() {
			continue
		}

		raw, err := os.ReadFile(skillFile)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", skillFile, err)
		}
		fm, _, err := parseFrontmatter(raw)
		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", skillFile, err)
		}
		desc := fm["description"]
		if desc == "" {
			return nil, fmt.Errorf("missing description in %s", skillFile)
		}

		skills = append(skills, Skill{
			Name:        entry.Name(),
			Description: desc,
			Dir:         skillDir,
		})
	}

	sort.Slice(skills, func(i, j int) bool {
		return skills[i].Name < skills[j].Name
	})
	return skills, nil
}
