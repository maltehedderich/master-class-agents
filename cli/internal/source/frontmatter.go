package source

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// parseFrontmatter splits a markdown file with leading YAML frontmatter into
// a flat string map of frontmatter values and the post-frontmatter body.
//
// Only top-level scalar keys are returned. Values may be unquoted, single
// quoted, or double quoted; values are returned with surrounding quotes
// stripped. List or nested values keep their raw rhs text.
//
// The frontmatter must start on line 1 with `---` and end with a closing
// `---`. Files without a leading frontmatter return an error.
func parseFrontmatter(data []byte) (map[string]string, string, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	if !scanner.Scan() {
		return nil, "", errors.New("empty file")
	}
	if scanner.Text() != "---" {
		return nil, "", errors.New("file does not start with frontmatter delimiter '---'")
	}

	fm := make(map[string]string)
	closed := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			closed = true
			break
		}
		key, value, ok := splitKeyValue(line)
		if !ok {
			continue
		}
		fm[key] = value
	}
	if !closed {
		return nil, "", errors.New("frontmatter not closed with '---'")
	}

	var body bytes.Buffer
	for scanner.Scan() {
		body.WriteString(scanner.Text())
		body.WriteByte('\n')
	}
	if err := scanner.Err(); err != nil {
		return nil, "", fmt.Errorf("scan: %w", err)
	}

	return fm, body.String(), nil
}

func splitKeyValue(line string) (string, string, bool) {
	idx := strings.Index(line, ":")
	if idx <= 0 {
		return "", "", false
	}
	key := strings.TrimSpace(line[:idx])
	if key == "" {
		return "", "", false
	}
	// Skip indented continuation lines and comment lines.
	if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") || strings.HasPrefix(key, "#") {
		return "", "", false
	}
	value := strings.TrimSpace(line[idx+1:])
	value = stripQuotes(value)
	return key, value, true
}

func stripQuotes(s string) string {
	if len(s) >= 2 {
		first, last := s[0], s[len(s)-1]
		if (first == '"' && last == '"') || (first == '\'' && last == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
