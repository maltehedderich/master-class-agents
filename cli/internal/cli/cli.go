// Package cli wires the cobra commands for the mcagents binary.
//
// Run is the testable entry point. main() calls Execute, which delegates to
// Run with os.Args, os.Stdin, os.Stdout, os.Stderr.
package cli

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/maltehedderich/master-class-agents/cli/internal/fsutil"
)

// exit codes returned by Run.
const (
	exitOK              = 0
	exitEnvironment     = 1
	exitInput           = 2
	exitArtifactFailure = 3
	exitInterrupted     = 130
)

// usageError marks errors triggered by malformed user input. Run translates
// them to exit code 2.
type usageError struct{ err error }

func (u *usageError) Error() string { return u.err.Error() }
func (u *usageError) Unwrap() error { return u.err }

func newUsageError(format string, a ...any) error {
	return &usageError{err: fmt.Errorf(format, a...)}
}

// artifactFailureError wraps the per-file errors collected by the install
// loop. Run translates it to exit code 3.
type artifactFailureError struct{ msg string }

func (a *artifactFailureError) Error() string { return a.msg }

// Execute runs the root command using the process's stdio. It is the entry
// point used by main(). Returns nil on success and a non-nil error otherwise;
// callers are responsible for printing and choosing an exit code, but
// mcagents/main.go simply prints and exits 1 for backward compatibility -
// real exit code routing happens in Run.
func Execute(version string) error {
	code := Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	if code != exitOK {
		os.Exit(code)
	}
	return nil
}

// Run is the testable entry. It builds a fresh root cobra command, attaches
// the provided IO, runs it, and translates errors to exit codes.
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	root := newRootCmd(stdin, stdout, stderr)
	root.SetArgs(args)
	root.SetIn(stdin)
	root.SetOut(stdout)
	root.SetErr(stderr)
	root.SilenceUsage = true
	root.SilenceErrors = true

	err := root.Execute()
	if err == nil {
		return exitOK
	}
	return classifyError(err, stderr)
}

func classifyError(err error, stderr io.Writer) int {
	var ue *usageError
	if errors.As(err, &ue) {
		_, _ = fmt.Fprintln(stderr, "error:", ue.Error())
		return exitInput
	}
	var af *artifactFailureError
	if errors.As(err, &af) {
		_, _ = fmt.Fprintln(stderr, af.Error())
		return exitArtifactFailure
	}
	if errors.Is(err, fsutil.ErrAborted) {
		_, _ = fmt.Fprintln(stderr, "aborted")
		return exitInterrupted
	}
	_, _ = fmt.Fprintln(stderr, "error:", err.Error())
	return exitEnvironment
}

func newRootCmd(stdin io.Reader, stdout, stderr io.Writer) *cobra.Command {
	root := &cobra.Command{
		Use:   "mcagents",
		Short: "Install master-class-agents agents and skills into your tool of choice",
	}
	root.AddCommand(newInstallCmd(stdin, stdout, stderr))
	root.AddCommand(newListCmd(stdout, stderr))
	root.Version = "" // overridden by Execute via SetVersionTemplate; left for tests
	return root
}
