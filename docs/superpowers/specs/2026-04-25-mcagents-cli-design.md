# Design: `mcagents` CLI

**Status:** Approved
**Date:** 2026-04-25
**Replaces:** `scripts/install-{claude,codex,copilot,gemini}-agents.sh`

## Summary

Build a Go CLI named `mcagents` that installs agents and skills from this repository into one of four target tools — Claude Code, Codex, GitHub Copilot, or Gemini CLI. The CLI is hybrid: it runs as an interactive picker when invoked without flags, and runs non-interactively when `--tool` and selection flags are provided. It replaces all four existing shell installers and adds first-class skill installation, which is currently a manual copy step.

## Goals

- One install command per user, regardless of target tool.
- Selective install: pick any subset of agents and skills, instead of "all or nothing".
- First-class support for skills, not only agents.
- Non-interactive flag mode for CI and dotfiles automation.
- Interactive picker for first-time users and ad-hoc installs.
- Zero Go install required after `git clone` (bootstrap wrapper handles it).

## Non-goals (v1)

- Uninstall, update, or sync subcommands. (Re-running `install` is the update path; `rm` is the uninstall path.)
- Network fetch of agents/skills from a remote registry. v1 reads from a local clone.
- Embedded artifacts (Go `embed`). v1 always reads from the repo on disk.
- Concurrent-install safety. Two `mcagents install` runs at once is user error.
- Backups before overwrite. The conflict prompt's `[d]iff` covers the inspection need.

## Distribution

Three install paths, all supported:

1. **Repo bootstrap (primary).** Users `git clone` and run `./mcagents install`. The repo-root `./mcagents` shell wrapper:
   - Looks for `cli/bin/<os>-<arch>/mcagents` (gitignored).
   - If absent, downloads the matching binary from the latest GitHub Release into that path.
   - Falls back to `go run ./cli/cmd/mcagents` if neither a cached binary nor a release is available, and `go` is on `$PATH`.
2. **`go install`.** `go install github.com/maltehedderich/master-class-agents/cli/cmd/mcagents@latest`. Requires Go; produces a binary on `$PATH`. Users still need a clone (or `--repo` pointing at one) for source content.
3. **GitHub Releases.** GoReleaser builds binaries for darwin-arm64, darwin-amd64, linux-arm64, linux-amd64, windows-amd64. These power path #1 and are also downloadable directly.

`cli/bin/` is `.gitignore`d. Binaries are never committed to the repo.

## Source resolution

The CLI must locate the repo's `agents/` and `skills/` directories before it can do anything. Resolution order:

1. `--repo PATH` flag.
2. `MCAGENTS_REPO` environment variable.
3. Walk up from `cwd` until both `agents/` and `skills/` exist as siblings.

If none of these resolve, exit with code 1 and a clear message:

> cannot find master-class-agents repo (looked for agents/ and skills/ from cwd up to /). Run from a clone, or pass --repo PATH or set MCAGENTS_REPO.

## Command surface

```
mcagents install
  --tool {claude|codex|copilot|gemini}    Required when stdin is not a TTY
  --agents <name,name,...>                "all" selects every agent; omit to skip agents
  --skills <name,name,...>                same semantics as --agents
  --dest <path>                           Override the tool's default destination
  --force                                 Overwrite without prompting
  --no-clobber                            Skip files that already exist (mutually exclusive with --force)
  --repo <path>                           Override source repo location
  --dry-run                               Print the plan; write nothing
  -v, --verbose                           Log every file action

mcagents list
  --tool {claude|codex|copilot|gemini}    Show destination column for that tool
  --json                                  Emit JSON instead of a table

mcagents --version
mcagents --help
```

Tool names are canonical: `claude`, `codex`, `copilot`, `gemini`. They match the tools' own CLI start commands. No aliases.

The picker runs when `mcagents install` is invoked with neither `--agents` nor `--skills`. If `--tool` is missing in that case and stdin is a TTY, the picker covers tool selection too. If stdin is not a TTY and `--tool` is missing, exit 2 with a usage error.

## Architecture

```
master-class-agents/
  mcagents                               Bootstrap wrapper (sh, executable, ~40 lines)
  agents/                                Source agents (unchanged)
  skills/                                Source skills (unchanged)
  guides/                                (unchanged)
  .gitignore                             + cli/bin/
  cli/
    go.mod                               Module: github.com/maltehedderich/master-class-agents/cli
    .goreleaser.yaml
    bin/                                 .gitignored, holds downloaded/built binaries
    cmd/mcagents/main.go                 Entry: Cobra root, version, wires subcommands
    internal/cli/
      install.go                         install subcommand
      list.go                            list subcommand
      picker.go                          huh-based interactive selector
      conflict.go                        Conflict prompt closure
    internal/source/
      source.go                          Resolve, load agents/skills
      agent.go                           Parse agents/*.agent.md
      skill.go                           Parse skills/<name>/SKILL.md
      testdata/                          Fixture repo for tests
    internal/installer/
      installer.go                       Installer interface + Registry
      claude.go                          Rewrites name frontmatter
      codex.go                           Converts agents to native Codex TOML files
      copilot.go                         Copies as-is
      gemini.go                          Wraps body in SKILL.md folder
    internal/fsutil/
      write.go                           Atomic write + conflict resolution
  .github/workflows/
    cli.yml                              Test + lint on push/PR (paths: cli/**)
    release.yml                          GoReleaser on tag
```

### Dependencies

- `github.com/spf13/cobra` — command parsing.
- `github.com/charmbracelet/huh` — interactive form picker (sits on Bubble Tea).
- `github.com/sergi/go-diff/diffmatchpatch` — unified diff for the conflict prompt's `[d]iff` option.
- Stdlib for everything else (file I/O, frontmatter parsing — small enough to write inline).

No cgo. Single static binary on every platform.

## Core interfaces

### `internal/source`

```go
type Source struct {
    Root string
}

type Agent struct {
    Name        string            // "backend-engineer", from filename without .agent.md
    Description string            // from frontmatter
    Path        string            // absolute path to agents/<name>.agent.md
    Frontmatter map[string]string // parsed YAML frontmatter (string-only values)
    Body        string            // post-frontmatter content
}

type Skill struct {
    Name        string // "privacy-policy", from skills/<name>/
    Description string // from skills/<name>/SKILL.md frontmatter
    Dir         string // absolute path to skills/<name>/
}

func Resolve(repoFlag string) (*Source, error)
func (s *Source) Agents() ([]Agent, error)
func (s *Source) Skills() ([]Skill, error)
```

### `internal/installer`

```go
type Kind int
const (KindAgent Kind = iota; KindSkill)

type Action int
const (ActionOverwrite Action = iota; ActionSkip; ActionAbort)

type Options struct {
    Force      bool
    NoClobber  bool
    DryRun     bool
    OnConflict func(path string, existing, incoming []byte) Action
}

type Result struct {
    Written []string
    Skipped []string
}

type Installer interface {
    Name() string                       // "claude", "codex", "copilot", "gemini"
    DefaultDir(kind Kind) string        // ~/.claude/agents, ~/.codex/agents, etc.
    InstallAgent(a source.Agent, destRoot string, opts Options) (Result, error)
    InstallSkill(s source.Skill, destRoot string, opts Options) (Result, error)
}

type Registry struct{ /* ... */ }
func (r *Registry) Get(name string) (Installer, bool)
func DefaultRegistry() *Registry
```

Each tool gets its own file (`claude.go`, `codex.go`, `copilot.go`, `gemini.go`) implementing the interface. Per-tool transforms:

| Tool    | Agent transform                                                                                                                                        | Skill transform                                       |
| ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------- |
| Copilot | Copy as-is to `<dest>/<name>.agent.md`                                                                                                                  | Copy `skills/<name>/` verbatim to `<dest>/<name>/`    |
| Claude  | Rewrite frontmatter `name:` to `"master-class-agents:<name>"`; remove source `tools:` so Claude inherits all available tools; output to `<dest>/<name>.agent.md` | Copy `skills/<name>/` verbatim to `<dest>/<name>/`    |
| Codex   | Strip the source frontmatter; emit native custom-agent TOML with `name`, `description`, and `developer_instructions`; write to `<dest>/<name>.toml` | Copy `skills/<name>/` verbatim to `<dest>/<name>/`    |
| Gemini  | Strip the source frontmatter; emit a new frontmatter block with `name: "<name>"` and `description: "<description>"`; write to `<dest>/<name>/SKILL.md` | Copy `skills/<name>/` verbatim to `<dest>/<name>/`    |

Default destinations:

| Tool    | Agents               | Skills               |
| ------- | -------------------- | -------------------- |
| Copilot | `~/.copilot/agents`  | `~/.copilot/skills`  |
| Claude  | `~/.claude/agents`   | `~/.claude/skills`   |
| Codex   | `~/.codex/agents`    | `~/.codex/skills`    |
| Gemini  | `~/.gemini/skills`   | `~/.gemini/skills`   |

(Codex installs native custom agents. Gemini installs agents-as-skills, so both Gemini kinds land under its `skills/` directory. All defaults are overridable via `--dest`.)

### `internal/fsutil`

`WriteFile(dest string, content []byte, opts Options) (Action, error)` does the atomic write:
1. If `dest` exists and `bytes.Equal(existing, content)`: return `ActionSkip`, no write.
2. If `dest` exists and differs:
   - `opts.NoClobber` → return `ActionSkip`.
   - `opts.Force` → fall through to write.
   - `opts.OnConflict` set → call it, act on the returned `Action`.
   - Otherwise → write (default for non-interactive flag mode).
3. Write to `<dest>.tmp`, `fsync`, `os.Rename` to `dest`. On rename failure, `os.Remove` the temp file before returning the error.

### `internal/cli/conflict.go`

The interactive `OnConflict` is a closure that holds two booleans across the run: `overwriteAll` and `skipAll`. After the user picks `[A]` or `[N]`, subsequent conflicts auto-resolve without prompting. The `[d]iff` option renders a unified diff via `github.com/sergi/go-diff/diffmatchpatch`.

## Data flow

### `mcagents install` (interactive)

1. Resolve source.
2. Load catalog (agents + skills).
3. Run picker: tool → agents → skills → confirm summary.
4. Resolve installer + destinations.
5. Build `Options` with the interactive `OnConflict` closure.
6. Loop selected agents and skills through the installer.
7. Print summary.

### `mcagents install --tool claude --agents all`

1. Resolve source.
2. Load catalog.
3. Skip picker.
4. Resolve installer + destinations (`--dest` overrides `DefaultDir` if set).
5. Build `Options` from flags. `OnConflict` is `nil` (default behaviour applies).
6. Install loop.
7. Plain-text summary on stdout.

### `mcagents list --tool claude`

1. Resolve source.
2. Print a table with columns: `NAME | TYPE | DESCRIPTION | DEST`. Without `--tool`, omit the `DEST` column.
3. With `--json`, emit `[{name, type, description, dest?}, ...]`.

## Error handling

Three classes:

- **Input errors** (unknown tool, unknown agent, conflicting flags) → exit 2, no stack trace.
- **Environment errors** (no source repo, dest not writable, malformed frontmatter) → exit 1, clear message.
- **Per-file errors** → collected, do not abort the run, summarized at the end. Exit 3 if any artifact failed.

Ctrl-C in the picker exits cleanly via huh. Ctrl-C during the install loop is caught with `signal.NotifyContext`; the in-flight atomic write finishes its rename (or fails cleanly, leaving only the temp file which is then removed), then the loop exits with code 130. Files already written stay written; the next artifact in the loop is not started.

Verbose logging is opt-in (`-v`); default output is one line per artifact plus a summary.

## Testing

### Unit tests — `internal/installer/*_test.go`

One file per installer. Each runs against `t.TempDir()` and asserts exact file contents.

- `claude_test.go`: name rewrite covers `name: foo`, `name: "foo"`, missing `name`, multi-line frontmatter.
- `codex_test.go`: body extraction strips original frontmatter and emits native custom-agent TOML; rejects agent names outside `[a-z0-9-]`.
- `gemini_test.go`: body extraction strips original frontmatter and rewraps; rejects skill names outside `[a-z0-9-]`.
- `copilot_test.go`: byte-for-byte copy.

### Source-loader tests — `internal/source/*_test.go`

- Frontmatter parser handles quoted values, unquoted values, descriptions with colons.
- `Resolve` walks up correctly; clear error when no clone is found.
- Missing `description` produces an error pointing at the file.

Fixtures live under `internal/source/testdata/` so tests don't depend on real `agents/`/`skills/` content.

### CLI integration tests — `internal/cli/*_test.go`

Drive `mcagents install` end-to-end with flags, pointing `--repo` at a fixture clone and `--dest` at `t.TempDir()`. Cover:

- Single agent, single tool → expected file at expected path.
- `--agents all` → every fixture agent installed.
- `--no-clobber` → existing differing file untouched.
- `--force` → overwrites without prompting.
- `--dry-run` → no writes; plan printed.
- Unknown `--tool` / unknown agent → exit 2 with the expected message.

Picker tests are minimal: one happy-path test using `teatest`. We don't exhaustively test TUI rendering.

### Fixture repo

`cli/internal/source/testdata/fixture-repo/` with two minimal artifacts:
```
agents/sample-agent.agent.md
skills/sample-skill/SKILL.md
```

### What we don't test in CI

- The bootstrap shell wrapper (`./mcagents`). Manual smoke test on macOS + Linux per release.
- Real installs into `~/.claude/`, etc. Tests always use a temp dir.

### CI

`.github/workflows/cli.yml` runs on `push` and `pull_request` for paths under `cli/**`:
- `go test ./...` on `ubuntu-latest` and `macos-latest`.
- `go vet ./...`.
- `golangci-lint run`.
- `goreleaser check`.

`.github/workflows/release.yml` runs on tag push, executes GoReleaser, attaches binaries to the GitHub Release.

## Migration

In the same change set:

- Add `cli/` module and bootstrap wrapper `./mcagents`.
- Delete `scripts/install-{claude,codex,copilot,gemini}-agents.sh`.
- Update `README.md`:
  - Replace the four-script install table with a single `./mcagents install` snippet plus `go install` and release-binary alternatives.
  - Remove the "skills are installed manually" caveat.
- Add `cli/bin/` to `.gitignore`.

No backwards compatibility shims for the old scripts. They're deleted in the same PR; the README points everyone at the new flow.

## Out of scope, but worth flagging for v2

- `mcagents uninstall <tool>` / `--all`.
- `mcagents doctor` to diagnose dest-dir permissions, conflicting installs.
- Embedded artifacts so `go install` works without a clone.
- Per-tool config file (`.mcagents.yaml`) for repeatable installs.
