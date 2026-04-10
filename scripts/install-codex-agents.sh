#!/bin/sh

set -eu

usage() {
  cat <<'EOF'
Usage: install-codex-agents.sh [TARGET_AGENTS_DIR]

Installs Codex agent role files generated from this repository's ./agents folder.

Install target resolution:
1. Positional TARGET_AGENTS_DIR argument
2. CODEX_AGENTS_DIR environment variable
3. CODEX_HOME/agents environment variable
4. $HOME/.codex/agents
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
TARGET_DIR="${1:-${CODEX_AGENTS_DIR:-$CODEX_ROOT/agents}}"

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

extract_body() {
  file_path=$1

  awk '
    NR == 1 && $0 == "---" {
      in_frontmatter = 1
      next
    }

    in_frontmatter && $0 == "---" {
      in_frontmatter = 0
      just_closed = 1
      next
    }

    just_closed {
      just_closed = 0
      if ($0 == "") {
        next
      }
    }

    !in_frontmatter {
      print
    }
  ' "$file_path"
}

toml_escape() {
  printf '%s' "$1" | sed 's/\\/\\\\/g; s/"/\\"/g'
}

if [ ! -d "$SOURCE_DIR" ]; then
  printf 'Error: source directory not found: %s\n' "$SOURCE_DIR" >&2
  exit 1
fi

mkdir -p "$TARGET_DIR"

installed_count=0

for source_file in "$SOURCE_DIR"/*.agent.md; do
  [ -e "$source_file" ] || continue

  role_name=$(basename "$source_file" .agent.md | tr '-' '_')
  description=$(extract_frontmatter_value description "$source_file" || true)
  description=$(strip_wrapping_quotes "$description")

  if [ -z "$description" ]; then
    printf 'Error: missing description in %s\n' "$source_file" >&2
    exit 1
  fi

  body=$(extract_body "$source_file")

  if [ -z "$body" ]; then
    printf 'Error: missing developer instructions in %s\n' "$source_file" >&2
    exit 1
  fi

  if printf '%s\n' "$body" | grep -q "'''"; then
    printf "Error: cannot convert %s because it contains triple single quotes\n" "$source_file" >&2
    exit 1
  fi

  case "$role_name" in
    ''|*[!abcdefghijklmnopqrstuvwxyz0123456789_]*)
      printf 'Error: unsupported Codex role name derived from %s: %s\n' "$source_file" "$role_name" >&2
      exit 1
      ;;
  esac

  target_file="$TARGET_DIR/$role_name.toml"

  {
    printf '# Generated from %s\n' "$(basename "$source_file")"
    printf 'name = "%s"\n' "$(toml_escape "$role_name")"
    printf 'description = "%s"\n' "$(toml_escape "$description")"
    printf "developer_instructions = '''\n"
    printf '%s\n' "$body"
    printf "'''\n"
  } > "$target_file"

  installed_count=$((installed_count + 1))
  printf 'Installed %s as %s\n' "$(basename "$source_file")" "$(basename "$target_file")"
done

if [ "$installed_count" -eq 0 ]; then
  printf 'Error: no agent files found in %s\n' "$SOURCE_DIR" >&2
  exit 1
fi

printf 'Installed %s Codex agent(s) into %s\n' "$installed_count" "$TARGET_DIR"