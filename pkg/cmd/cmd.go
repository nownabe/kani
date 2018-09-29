package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return newCommand()
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kani",
		Short: "Kani is umai",
		Long: `
Kani is super good food.
Very very umai.`,
		Run: runHelp,
	}

	cmd.AddCommand(newCmdHelloWorld())

	return cmd
}

func runHelp(cmd *cobra.Command, asrs []string) {
	cmd.Help()
}

func newCmdHelloWorld() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "hello world",
		Run:   runHelloWorld,
	}
}

func runHelloWorld(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, world!")
}
