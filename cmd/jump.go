package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

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
				var err error
				var value string
				oflog.Init()
				err = algorithm.JumpGetInfo(value)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.JumpGetInfo!")
					return err
				}
				return nil
			},
		},
	},
}
