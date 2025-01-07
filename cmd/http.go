package cmd

import (
	"fmt"
	"oftools/algorithm"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var HttpCommand = &cli.Command{
	Name:    "http",
	Aliases: []string{"ht"},
	Usage:   "Greets the user with the provided name",
	Action: func(c *cli.Context) error {
		//var err error
		port := c.Args().Get(0)
		err := algorithm.Http(port)
		if err != nil {
			return fmt.Errorf("invalid decimal value: %v", err)
		}
		return nil
	},
}
