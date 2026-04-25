package source

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Agent is a parsed agents/<name>.agent.md artifact.
type Agent struct {
	Name        string
	Description string
	Path        string
	Frontmatter map[string]string
	Body        string
}

// Agents loads every agents/*.agent.md file under s.Root.
//
// Files are returned sorted by Name. Each file must have a frontmatter block
// with a non-empty description; missing descriptions return an error pointing
// at the file.
func (s *Source) Agents() ([]Agent, error) {
	dir := filepath.Join(s.Root, "agents")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read agents dir %s: %w", dir, err)
	}

	var agents []Agent
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".agent.md") {
			continue
		}
		path := filepath.Join(dir, name)
		agent, err := loadAgent(path)
		if err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}

	sort.Slice(agents, func(i, j int) bool {
		return agents[i].Name < agents[j].Name
	})
	return agents, nil
}

func loadAgent(path string) (Agent, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Agent{}, fmt.Errorf("read %s: %w", path, err)
	}

	fm, body, err := parseFrontmatter(raw)
	if err != nil {
		return Agent{}, fmt.Errorf("parse %s: %w", path, err)
	}

	desc := fm["description"]
	if desc == "" {
		return Agent{}, fmt.Errorf("missing description in %s", path)
	}

	name := strings.TrimSuffix(filepath.Base(path), ".agent.md")
	return Agent{
		Name:        name,
		Description: desc,
		Path:        path,
		Frontmatter: fm,
		Body:        body,
	}, nil
}
