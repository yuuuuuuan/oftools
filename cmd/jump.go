package cmd

import (
	"fmt"
	"oftools/algorithm"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var JumpCommand = &cli.Command{
	Name:    "jump",
	Aliases: []string{"j"},
	Usage:   "Greets the user with the provided name",
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "getinfo",
			Usage: "Convert a hex value to hexadecimal, binary, and octal",
			Action: func(c *cli.Context) error {
				//var err error
				err := algorithm.Jump()
				if err != nil {
					return fmt.Errorf("invalid decimal value: %v", err)
				}
				return nil
			},
		},
	},
}
