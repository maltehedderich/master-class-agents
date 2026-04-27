package installer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/maltehedderich/master-class-agents/cli/internal/source"
)

// Claude installs agents into Claude Code with a rewritten frontmatter
// `name:` field that namespaces the agent under "master-class-agents:<name>".
type Claude struct{}

// Name returns "claude".
func (*Claude) Name() string { return "claude" }

// DefaultDir returns ~/.claude/agents or ~/.claude/skills.
func (*Claude) DefaultDir(kind Kind) string {
	switch kind {
	case KindAgent:
		return homeJoin(".claude", "agents")
	default:
		return homeJoin(".claude", "skills")
	}
}

// InstallAgent rewrites the frontmatter `name:` to "master-class-agents:<a.Name>"
// (inserting one if missing), removes Copilot-specific tool aliases, and writes
// to destRoot/<name>.agent.md.
func (*Claude) InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error) {
	raw, err := os.ReadFile(a.Path)
	if err != nil {
		return Result{}, fmt.Errorf("read agent %s: %w", a.Path, err)
	}
	rewritten, err := rewriteClaudeName(raw, a.Name)
	if err != nil {
		return Result{}, fmt.Errorf("rewrite %s: %w", a.Path, err)
	}
	dest := filepath.Join(destRoot, a.Name+".agent.md")
	return writeOne(dest, rewritten, opts)
}

// InstallSkill recursively copies the source skill folder under destRoot.
func (*Claude) InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error) {
	return copySkillDir(s, destRoot, opts)
}

// rewriteClaudeName replaces the frontmatter `name:` line with the namespaced
// form, or inserts one before the closing `---` if missing. It also removes the
// source `tools:` field because the repository uses Copilot tool aliases, while
// Claude Code has its own tool names and safely inherits all tools when omitted.
func rewriteClaudeName(raw []byte, agentName string) ([]byte, error) {
	target := fmt.Sprintf("name: %q", "master-class-agents:"+agentName)

	scanner := bufio.NewScanner(bytes.NewReader(raw))
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var out bytes.Buffer
	state := 0 // 0=before fm, 1=in fm, 2=after fm
	wroteName := false
	skippingTools := false

	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case 0:
			if line == "---" {
				state = 1
				out.WriteString(line)
				out.WriteByte('\n')
				continue
			}
			out.WriteString(line)
			out.WriteByte('\n')
		case 1:
			if line == "---" {
				if !wroteName {
					out.WriteString(target)
					out.WriteByte('\n')
					wroteName = true
				}
				out.WriteString(line)
				out.WriteByte('\n')
				state = 2
				continue
			}
			if skippingTools {
				if isIndentedOrBlank(line) {
					continue
				}
				skippingTools = false
			}
			if isTopLevelKey(line, "name") {
				if !wroteName {
					out.WriteString(target)
					out.WriteByte('\n')
					wroteName = true
				}
				continue
			}
			if isTopLevelKey(line, "tools") {
				skippingTools = true
				continue
			}
			out.WriteString(line)
			out.WriteByte('\n')
		case 2:
			out.WriteString(line)
			out.WriteByte('\n')
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if state == 0 {
		return nil, fmt.Errorf("file does not start with frontmatter")
	}
	if state == 1 {
		return nil, fmt.Errorf("frontmatter not closed")
	}

	// Preserve absence of trailing newline in input if the original lacked one.
	result := out.Bytes()
	if len(raw) > 0 && raw[len(raw)-1] != '\n' && len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}
	return result, nil
}

// isTopLevelKey reports whether line declares the top-level YAML key `key`
// (i.e. starts at column 0, not indented).
func isTopLevelKey(line, key string) bool {
	if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
		return false
	}
	prefix := key + ":"
	if line == prefix {
		return true
	}
	if strings.HasPrefix(line, prefix) {
		next := line[len(prefix)]
		return next == ' ' || next == '\t'
	}
	return false
}

func isIndentedOrBlank(line string) bool {
	return strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") || strings.TrimSpace(line) == ""
}
