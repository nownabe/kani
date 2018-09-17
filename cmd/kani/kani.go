package main

import (
	"fmt"
	"os"

	"github.com/nownabe/kani/pkg/cmd"
)

func main() {
	command := cmd.New()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
