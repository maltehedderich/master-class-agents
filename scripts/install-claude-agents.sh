#!/bin/sh

set -eu

usage() {
  cat <<'EOF'
Usage: install-claude-agents.sh [TARGET_AGENTS_DIR]

Installs Claude Code agent files from this repository's ./agents folder.

Install target resolution:
1. Positional TARGET_AGENTS_DIR argument
2. CLAUDE_AGENTS_DIR environment variable
3. CLAUDE_DIR/agents environment variable
4. $HOME/.claude/agents
EOF
}

case "${1:-}" in
  -h|--help)
    usage
    exit 0
    ;;
esac

if [ "$#" -gt 1 ]; then
  usage >&2
  exit 1
fi

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
REPO_ROOT=$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)
SOURCE_DIR="$REPO_ROOT/agents"
CLAUDE_HOME="${CLAUDE_DIR:-$HOME/.claude}"
TARGET_DIR="${1:-${CLAUDE_AGENTS_DIR:-$CLAUDE_HOME/agents}}"

rewrite_claude_agent_name() {
  agent_name=$1
  file_path=$2

  awk -v agent_name="$agent_name" '
    BEGIN {
      claude_name = "master-class-agents:" agent_name
    }

    NR == 1 && $0 == "---" {
      in_frontmatter = 1
      found_frontmatter = 1
      print
      next
    }

    in_frontmatter && $0 == "---" {
      if (!rewrote_name) {
        print "name: \"" claude_name "\""
      }

      print
      in_frontmatter = 0
      next
    }

    in_frontmatter && $0 ~ /^name:[[:space:]]*/ {
      print "name: \"" claude_name "\""
      rewrote_name = 1
      next
    }

    {
      print
    }

    END {
      if (!found_frontmatter) {
        exit 1
      }
    }
  ' "$file_path"
}

if [ ! -d "$SOURCE_DIR" ]; then
  printf 'Error: source directory not found: %s\n' "$SOURCE_DIR" >&2
  exit 1
fi

mkdir -p "$TARGET_DIR"

installed_count=0

for source_file in "$SOURCE_DIR"/*.agent.md; do
  [ -e "$source_file" ] || continue

  agent_name=$(basename "$source_file" .agent.md)
  target_file="$TARGET_DIR/$(basename "$source_file")"
  temp_file="$TARGET_DIR/.${agent_name}.agent.md.tmp"

  if ! rewrite_claude_agent_name "$agent_name" "$source_file" > "$temp_file"; then
    rm -f "$temp_file"
    printf 'Error: failed to rewrite Claude agent name in %s\n' "$source_file" >&2
    exit 1
  fi

  mv "$temp_file" "$target_file"

  installed_count=$((installed_count + 1))
  printf 'Installed %s as %s\n' "$(basename "$source_file")" "master-class-agents:$agent_name"
done

if [ "$installed_count" -eq 0 ]; then
  printf 'Error: no agent files found in %s\n' "$SOURCE_DIR" >&2
  exit 1
fi

printf 'Installed %s Claude Code agent(s) into %s\n' "$installed_count" "$TARGET_DIR"
