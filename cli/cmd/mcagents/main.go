package main

import (
	"fmt"
	"os"

	"github.com/maltehedderich/master-class-agents/cli/internal/cli"
)

// version is overridden at build time via -ldflags "-X main.version=...".
var version = "dev"

func main() {
	if err := cli.Execute(version); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
