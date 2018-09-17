package cmd

import "fmt"

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c *Command) Execute() error {
	fmt.Println("Hello, world!")
	return nil
}
