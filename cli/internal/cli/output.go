package cli

import (
	"fmt"
	"io"
)

func writeLine(w io.Writer, a ...any) error {
	_, err := fmt.Fprintln(w, a...)
	return err
}

func writef(w io.Writer, format string, a ...any) error {
	_, err := fmt.Fprintf(w, format, a...)
	return err
}
