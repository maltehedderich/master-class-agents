---
name: idiomatic-go
description: "Write, review, or refactor Go code against the standard library and masterclass principles. Use when designing a small Go service or package, reviewing a PR for idiomatic Go, refactoring non-idiomatic code (deep hierarchies, premature interfaces, framework imports, leaked goroutines), debating package layout, designing error flow with wrapping and behavioral checks, sizing interfaces on the consumer side, deciding when concurrency is actually needed, or wiring up the Go toolchain (gofmt, go vet, golangci-lint, go test -race -cover, pprof)."
argument-hint: "Goal (write/review/refactor), the Go code or brief, target package boundary, runtime context (CLI, HTTP service, library), and any constraints (Go version, allowed dependencies, performance budget)"
---

# Idiomatic Go

## What This Skill Does

Turn a brief, a draft, or a piece of non-idiomatic Go into code that reads like the standard library: small interfaces on the consumer side, explicit error flow, useful zero values, and concurrency only where it earns its keep.

This skill treats Go as a constrained design target with one controlling rule: simplicity is a feature. Boring Go survives a decade of change in production. Clever Go costs more than it saves.

Use this skill to:

- Write a new package, CLI, or HTTP service in idiomatic Go.
- Review or critique a Go PR, file, or function against the masterclass principles.
- Refactor non-idiomatic Go (deep hierarchies, "manager" structs, preemptive interfaces, leaked goroutines, framework imports) into something composable.
- Decide whether a problem actually needs concurrency, generics, or a third-party dependency before reaching for them.
- Wire up the toolchain (`gofmt`, `go vet`, `golangci-lint`, `go test -race -cover`, `pprof`) so the language's authority is actually applied.

This skill is for Go craft. It is not for ranking Go against other languages, designing a multi-service architecture, picking a cloud provider, or writing project-specific style guides.

## When to Use

- The user has Go code (a function, file, or PR) and wants it shaped into something idiomatic.
- The user has a brief ("build a small ingest service in Go") and wants the first cut written the way the standard library would write it.
- An existing Go service feels Java-shaped — deep packages, getters/setters, framework configuration, service/manager classes — and needs flattening.
- A user is reaching for goroutines, channels, generics, or a new dependency and wants a sanity check before the code lands.
- A team is setting up a Go repo and needs the toolchain wired into CI before the second feature.

Do not use this skill for:

- Multi-service architecture, infra design, or cloud provider selection (use the backend or SRE agents).
- Language comparisons or "should we use Go" decisions.
- Mechanical formatting work that `gofmt` already enforces.
- Very large refactors that need a real plan and review checkpoints rather than a single skill pass.

## Inputs To Gather

Ask only for what is missing. Do not block on optional inputs.

- `goal` — `write`, `review`, `refactor`, or `toolchain`.
- `code_or_brief` — the Go code to review or refactor, or a short brief for new code.
- `runtime_context` — CLI, HTTP service, library, batch job, or test helper. Different shapes have different defaults (handlers, `context` propagation, `io.Reader`/`io.Writer` boundaries).
- `package_boundary` — what this package is responsible for in one sentence. If you cannot say it, the boundary is wrong.
- `constraints` — Go version, allowed dependencies, performance budget, target deployment (single static binary, container, library import).

Helpful but optional:

- Existing project layout and naming conventions to align with.
- Known performance hotspots, profiling output, or `-race` failures.
- The single user-visible behavior the change is meant to support.

If the brief is "make this Go code better," ask what the package is for and which behavior matters before touching anything. If the input is a 500-line `manager.go`, ask which one responsibility you may extract first rather than rewriting wholesale.

## Non-Negotiable Constraints

Treat these as defaults unless the user explicitly overrides them.

1. **Clarity over cleverness.** If a function cannot be explained in one sentence, split it. Use `gofmt` formatting, idiomatic short names in narrow scopes (`i`, `r`, `buf`) and full names at package boundaries. The standard library — `io`, `net/http`, `sync`, `encoding/json` — is the style reference.
2. **Errors are values.** Return errors as the last value. Wrap with `fmt.Errorf("doing X: %w", err)` to preserve the chain. Use `errors.Is` / `errors.As` for behavioral checks. Never log-and-return — pick one. Reserve `panic` for impossible states, never for control flow.
3. **Accept interfaces, return concrete types — and keep interfaces small.** Interfaces belong on the consumer side, defined by what _that_ caller needs (1–3 methods). Do not declare interfaces preemptively "in case someone needs them." Return structs so callers see real capabilities. `io.Reader` / `io.Writer` are the model.
4. **Reach for goroutines last.** Start synchronous. Add concurrency only when there is a real reason. Every goroutine has a known lifetime, propagated `context.Context`, and a clear exit path. A `sync.Mutex` around a struct is often clearer than channels for shared state. Channels shine for ownership transfer and pipelines. Run tests with `-race`.
5. **Design for the zero value and for composition.** Pick struct fields and defaults so `var x T` works (`bytes.Buffer{}`, `sync.Mutex{}`). Prefer struct embedding for composition. Build behavior by combining small pieces (`http.Handler`, middleware, `io` interfaces) rather than configuring a giant object.
6. **The toolchain is part of the language.** `gofmt` is non-negotiable. `go vet` and `golangci-lint` run in CI. `go test -race -cover` is the default test invocation. Profile with `pprof` and the execution tracer before optimizing. Use modules with intentional, pinned versions.
7. **Be conservative with dependencies.** "A little copying is better than a little dependency." Lean on the standard library first — for an HTTP service often `net/http`, `encoding/json`, `database/sql`, and `log/slog` are enough. When you do add a dependency, prefer small, focused libraries with clear maintainers.
8. **Test behavior, not implementation.** Table-driven tests with `t.Run` subtests. Test exported behavior. Use `httptest.Server` and real I/O over fakes when feasible. Keep test data in `testdata/`. The small consumer-side interface (constraint 3) makes substitution trivial without a mocking framework.

## Procedure

Follow the steps in order. Do not skip the package-boundary step.

### Step 1 — Confirm goal, boundary, and runtime context

State, in one sentence each:

- What this code is responsible for.
- What runtime context it lives in (CLI, HTTP handler, library, batch job).
- What user-visible behavior the change is meant to enable or protect.

If any of those is unclear, ask one targeted question rather than designing or reviewing around ambiguity. A package without a one-sentence responsibility is the real bug.

### Step 2 — Read the existing code (or the standard library) before writing

For `review` or `refactor`:

- Read the whole file or package, not just the diff.
- Identify the consumer-side interfaces actually used. Note any interface declared on the producer side "for flexibility."
- Trace error paths from origin to where they are handled or surfaced.
- Identify every goroutine and ask: what cancels it, and what closes its channels?

For `write`:

- Skim the most relevant standard library package (`net/http`, `io`, `encoding/json`, `database/sql`, `log/slog`) for the shape this code should match.
- Look at one nearby package in the same repo for naming and layout conventions.

### Step 3 — Design the smallest API that does the job

Before writing or rewriting code:

- List the exported identifiers you intend to add. Justify each one.
- For each consumer, write the 1–3 method interface they actually need on _their_ side, not on the producing package.
- Define struct types so the zero value is useful where possible.
- Decide error types: do callers branch on behavior? If yes, define sentinel errors or typed errors and document them. If no, wrap and propagate.

If the API has more than ~5 exported identifiers per file, or an interface with more than 3 methods, treat it as a smell and revisit the boundary.

### Step 4 — Implement or refactor in idiomatic shape

Write or rewrite with these defaults:

- Functions take `context.Context` as the first parameter for anything that does I/O, blocks, or spawns goroutines.
- Errors are wrapped with `%w` and a short doing-X verb phrase. No stack traces, no log-and-return.
- Synchronous by default. Goroutines only with a clear lifetime and exit path.
- Embedding over inheritance-style hierarchies. No `BaseService`, `AbstractFoo`, or "manager" coupling.
- Structured logging via `log/slog` with stable keys. No `fmt.Println` for ops output.
- HTTP handlers are small `http.Handler` values, composed via middleware. No framework imports unless the constraints require one.
- For shared state, prefer a small struct guarded by a `sync.Mutex` over channel choreography.

### Step 5 — Wire up the toolchain checks

Before declaring code ready:

- `gofmt -l` returns no files. (`goimports -l` if used.)
- `go vet ./...` passes.
- `golangci-lint run` passes (or the chosen subset documented in the repo).
- `go test -race -cover ./...` passes. Coverage is honest — exported behavior, not glue.
- For performance-sensitive code, capture a `pprof` profile or a benchmark in `testing.B` before making optimization claims.

If any of these steps are not part of CI yet, add them before adding more features.

### Step 6 — Final pass for Go-specific anti-patterns

Before delivering, scan for:

- Interfaces declared on the producing side without two concrete implementations behind them.
- Goroutines started without a `context`, deferred `Close`, or termination signal.
- `panic` used for control flow, or `recover` outside a top-level boundary.
- Deep package hierarchies, `internal/util`, or "manager"/"service" classes used as namespaces.
- Configuration frameworks where a struct literal would do.
- `interface{}` or `any` where a small typed interface or generic would say what is meant.
- Tests that mock everything and verify implementation rather than behavior.

Cut, inline, or rewrite anything that fails this scan.

## Branching Logic

- **Writing a new HTTP service.** Default to `net/http` plus `log/slog`, `encoding/json`, and `database/sql`. Define handlers as small `http.Handler` values. Compose middleware. Propagate `context.Context`. Resist frameworks unless the constraints require one.
- **Writing a new library.** Optimize for the caller. Define types whose zero value is useful. Return concrete types. Let consumers declare their own interfaces. Keep the dependency footprint near zero.
- **Reviewing a PR.** Lead with the structural issues (package boundary, interface placement, goroutine lifetime). Then error flow. Then style. Do not let `gofmt`-level nits crowd out design feedback.
- **Refactoring "Java-shaped" Go.** Do not rewrite wholesale. Pick one responsibility, extract it into a flat package with a useful zero value, route one caller through it, and ship. Repeat.
- **Concurrency request.** Push back first. Most "needs concurrency" problems are I/O or design problems. If real, prefer pipelines with channels for ownership transfer or `sync.Mutex` for shared state. Always require `-race` and a documented exit path.
- **Generics request.** Confirm there are at least two concrete callers that share a type-parameterized shape. If not, write the concrete version and revisit later.
- **Dependency request.** Check the standard library first. If a dependency is justified, prefer small, focused libraries with active maintainers and stable APIs. Pin versions intentionally.
- **Toolchain-only task.** Wire `gofmt`, `go vet`, `golangci-lint`, and `go test -race -cover` into CI before adding linting rules. Add `pprof` and the execution tracer when there is a real performance question.

## Output Format

Match the deliverable to the task.

For `write`:

1. The Go code, formatted as `gofmt` would format it, in the file layout you recommend (one fenced block per file, with the path as a header).
2. A short **Notes** section: package boundary in one sentence, the consumer-side interfaces, what is intentionally _not_ included, and what the next change would be.
3. A **Toolchain** snippet showing how to run `gofmt`, `go vet`, `golangci-lint`, and `go test -race -cover` for this package.

For `review`:

1. A prioritized list of findings, grouped: structural / error flow / concurrency / style / tests. Each finding cites the file and rough location and explains the principle it violates.
2. A "leave it" section for things that look unidiomatic but are fine in this context — disagreement is allowed.
3. A short suggested patch (diff or before/after snippet) for the top one or two issues.

For `refactor`:

1. A one-paragraph summary: what responsibility is being extracted and why.
2. The refactored Go code in fenced blocks per file.
3. A migration note: which call sites change, in what order, and how to verify nothing regressed (tests, `-race` run, coverage on the new boundary).

For `toolchain`:

1. A minimal CI step (or `Makefile`/`justfile` snippet) that runs the standard toolchain.
2. A `golangci-lint` config that turns on a defensible default set, not everything.
3. One follow-up suggestion (benchmarks, pprof endpoint, execution tracer) if the project will outgrow the basics soon.

## Quality Checks

The work is ready only when all of the following are true:

1. The package or change can be described in one sentence, and that sentence matches the code.
2. Errors are values: wrapped with `%w`, checked with `errors.Is` / `errors.As` where behavior matters, and never both logged and returned.
3. Interfaces are small, declared on the consumer side, and justified by at least one real caller.
4. Concrete types are returned. Zero values are useful where reasonable.
5. Every goroutine has a known lifetime, a propagated `context.Context`, and a clear exit path. Tests pass under `-race`.
6. Composition is via small pieces (`http.Handler`, middleware, `io` interfaces, struct embedding), not via deep hierarchies or "manager" classes.
7. The standard library does the work it can. Each external dependency has a one-sentence justification.
8. `gofmt`, `go vet`, `golangci-lint`, and `go test -race -cover` all pass on the affected packages.
9. Tests exercise behavior, are table-driven where it helps, and use `httptest`/real I/O where realistic.
10. The output reads like code from `io`, `net/http`, or `sync` — boring, obvious, and easy to change.

## Failure Modes To Avoid

- Declaring interfaces on the producing side "in case someone needs them," then passing them everywhere.
- Goroutine and resource leaks: starting goroutines without `context`, deferred `Close`, or a clear termination signal.
- Java- or Python-shaped Go: deep package hierarchies, getters/setters, "manager"/"service"/"helper" classes, configuration frameworks, dependency-injection containers.
- Wrapping `net/http` in a framework before there is a reason to.
- Using `panic`/`recover` for control flow.
- Catching all errors at the top with `fmt.Errorf("error: %v", err)` instead of `%w`, losing the chain.
- Mocking every collaborator in tests so the test is a mirror of the implementation.
- Reaching for generics, channels, or new dependencies before two concrete callers justify them.
- Optimizing without a benchmark or `pprof` profile to point at.
- Treating `gofmt`/`go vet`/`golangci-lint`/`-race` as optional. They are part of the language.

## Default Stance

Boring Go is the goal. The standard library is the style guide. Small interfaces beat large ones. Useful zero values beat constructors. Synchronous beats concurrent until proven otherwise. One package with a clear sentence beats five with overlapping responsibilities. When in doubt, copy the shape of `io`, `net/http`, or `sync`, run `go test -race -cover`, and ship the simpler version.
