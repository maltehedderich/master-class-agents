#!/bin/sh

set -eu

usage() {
  cat <<'EOF'
Usage: install-codex-agents.sh [TARGET_SKILLS_DIR]

Installs Codex skills generated from this repository's ./agents folder.

Install target resolution:
1. Positional TARGET_SKILLS_DIR argument
2. CODEX_SKILLS_DIR environment variable
3. CODEX_HOME/skills environment variable
4. $HOME/.codex/skills

Legacy note:
- `CODEX_AGENTS_DIR` is still accepted for backward compatibility, but Codex now
  discovers custom instructions as skills under `~/.codex/skills`.
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
CODEX_ROOT="${CODEX_HOME:-$HOME/.codex}"
DEFAULT_TARGET_DIR="$CODEX_ROOT/skills"
TARGET_DIR="${1:-${CODEX_SKILLS_DIR:-${CODEX_AGENTS_DIR:-$DEFAULT_TARGET_DIR}}}"

extract_frontmatter_value() {
  key=$1
  file_path=$2

  awk -v key="$key" '
    NR == 1 && $0 == "---" {
      in_frontmatter = 1
      next
    }

    in_frontmatter && $0 == "---" {
      exit
    }

    in_frontmatter && $0 ~ ("^" key ":[[:space:]]*") {
      value = $0
      sub("^" key ":[[:space:]]*", "", value)
      print value
      exit
    }
  ' "$file_path"
}

strip_wrapping_quotes() {
  value=$1

  case "$value" in
    \"*\")
      value=${value#\"}
      value=${value%\"}
      ;;
    \'*\')
      value=${value#\'}
      value=${value%\'}
      ;;
  esac

  printf '%s\n' "$value"
}

yaml_escape() {
  printf '%s' "$1" | sed 's/\\/\\\\/g; s/"/\\"/g'
}

extract_body() {
  file_path=$1

  awk '
    NR == 1 && $0 == "---" {
      in_frontmatter = 1
      next
    }

    in_frontmatter && $0 == "---" {
      in_frontmatter = 0
      next
    }

    !in_frontmatter {
      print
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

  skill_name=$(basename "$source_file" .agent.md)
  description=$(extract_frontmatter_value description "$source_file" || true)
  description=$(strip_wrapping_quotes "$description")
  body=$(extract_body "$source_file")

  if [ -z "$description" ]; then
    printf 'Error: missing description in %s\n' "$source_file" >&2
    exit 1
  fi

  if [ -z "$body" ]; then
    printf 'Error: missing skill instructions in %s\n' "$source_file" >&2
    exit 1
  fi

  case "$skill_name" in
    ''|*[!abcdefghijklmnopqrstuvwxyz0123456789-]*)
      printf 'Error: unsupported Codex skill name derived from %s: %s\n' "$source_file" "$skill_name" >&2
      exit 1
      ;;
  esac

  skill_dir="$TARGET_DIR/$skill_name"
  target_file="$skill_dir/SKILL.md"

  mkdir -p "$skill_dir"

  {
    printf '%s\n' '---'
    printf 'name: "%s"\n' "$(yaml_escape "$skill_name")"
    printf 'description: "%s"\n' "$(yaml_escape "$description")"
    printf '%s\n' '---'
    printf '%s\n' "$body"
  } > "$target_file"

  installed_count=$((installed_count + 1))
  printf 'Installed %s\n' "$skill_name"
done

if [ "$installed_count" -eq 0 ]; then
  printf 'Error: no agent files found in %s\n' "$SOURCE_DIR" >&2
  exit 1
fi

printf 'Installed %s Codex skill(s) into %s\n' "$installed_count" "$TARGET_DIR"
printf 'Restart Codex to pick up new skills\n'
