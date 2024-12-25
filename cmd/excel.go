package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var ExcelCommand = &cli.Command{
	Name:    "excel",
	Aliases: []string{"e"},
	Usage:   "Greets the user with the provided name",
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "sum",
			Usage: "Convert a hex value to hexadecimal, binary, and octal",
			Action: func(ctx *cli.Context) error {
				
				return nil
			},
		},
	},
}
