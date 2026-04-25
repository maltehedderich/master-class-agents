package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/maltehedderich/master-class-agents/cli/internal/fsutil"
)

// newInteractiveConflictHandler returns a closure that prompts the user
// whenever an existing destination differs from the incoming bytes. It
// remembers an "all" choice for the rest of the run.
func newInteractiveConflictHandler(stdin io.Reader, stdout io.Writer) func(path string, existing, incoming []byte) fsutil.Action {
	overwriteAll := false
	skipAll := false
	reader := bufio.NewReader(stdin)

	return func(path string, existing, incoming []byte) fsutil.Action {
		if overwriteAll {
			return fsutil.ActionOverwrite
		}
		if skipAll {
			return fsutil.ActionSkip
		}

		for {
			fmt.Fprintf(stdout, "%s already exists with different content. [o]verwrite / [s]kip / [A]ll / [N]one / [d]iff / [q]uit? ", path)
			line, err := reader.ReadString('\n')
			if err != nil {
				return fsutil.ActionAbort
			}
			choice := strings.TrimSpace(strings.ToLower(line))
			switch choice {
			case "o", "overwrite":
				return fsutil.ActionOverwrite
			case "s", "skip":
				return fsutil.ActionSkip
			case "a", "all", "yes-all":
				overwriteAll = true
				return fsutil.ActionOverwrite
			case "n", "none", "skip-all":
				skipAll = true
				return fsutil.ActionSkip
			case "d", "diff":
				printDiff(stdout, existing, incoming)
				continue
			case "q", "quit", "abort":
				return fsutil.ActionAbort
			default:
				fmt.Fprintln(stdout, "  please pick one of o, s, A, N, d, q")
			}
		}
	}
}

func printDiff(w io.Writer, existing, incoming []byte) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(existing), string(incoming), false)
	fmt.Fprintln(w, dmp.DiffPrettyText(diffs))
}
