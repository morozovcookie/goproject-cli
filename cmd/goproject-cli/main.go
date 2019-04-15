package main

import (
	"fmt"
	"os"

	"github.com/morozovcookie/goproject-cli/pkg/cmd"
)

func main() {
	command := cmd.NewDefaultCommand()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
