#!/bin/sh

set -eu

usage() {
  cat <<'EOF'
Usage: install-copilot-agents.sh [TARGET_AGENTS_DIR]

Installs GitHub Copilot agent files from this repository's ./agents folder.

Install target resolution:
1. Positional TARGET_AGENTS_DIR argument
2. COPILOT_AGENTS_DIR environment variable
3. COPILOT_DIR/agents environment variable
4. $HOME/.copilot/agents

Tip: pass TARGET_AGENTS_DIR explicitly if you want to install into a
profile-specific VS Code agents folder instead of the default user-level one.
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
COPILOT_HOME="${COPILOT_DIR:-$HOME/.copilot}"
TARGET_DIR="${1:-${COPILOT_AGENTS_DIR:-$COPILOT_HOME/agents}}"

if [ ! -d "$SOURCE_DIR" ]; then
  printf 'Error: source directory not found: %s\n' "$SOURCE_DIR" >&2
  exit 1
fi

mkdir -p "$TARGET_DIR"

installed_count=0

for source_file in "$SOURCE_DIR"/*.agent.md; do
  [ -e "$source_file" ] || continue

  target_file="$TARGET_DIR/$(basename "$source_file")"
  cp "$source_file" "$target_file"
  installed_count=$((installed_count + 1))
  printf 'Installed %s\n' "$(basename "$source_file")"
done

if [ "$installed_count" -eq 0 ]; then
  printf 'Error: no agent files found in %s\n' "$SOURCE_DIR" >&2
  exit 1
fi

printf 'Installed %s GitHub Copilot agent(s) into %s\n' "$installed_count" "$TARGET_DIR"