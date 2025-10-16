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
	Usage:   "💖 Auto logging in jumpserver",
	HideHelpCommand:      true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "update",
			Usage: "Update your token",
			Action: func(c *cli.Context) error {
				var err error
				//var value string
				oflog.Init()
				err = algorithm.JumpUpdateToken()
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.JumpGetInfo!")
					return err
				}
				return nil
			},
		},

		// Hex to other systems
		{
			Name:  "server",
			Usage: "Open your jump server",
			Action: func(c *cli.Context) error {
				var err error
				//var value string
				oflog.Init()
				err = algorithm.JumpServer()
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.JumpGetInfo!")
					return err
				}
				return nil
			},
		},
	},
}
