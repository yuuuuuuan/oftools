package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var IworkCommand = &cli.Command{
	Name:            "iwork",
	Aliases:         []string{"i"},
	Usage:           "🍀 Greets the user with the provided name",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "get",
			Usage: "Convert a hex value to hexadecimal, binary, and octal",
			Action: func(c *cli.Context) error {
				var err error
				//var value string
				oflog.Init()
				err = algorithm.IworkGet()
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.IworkGet!")
					return err
				}
				return nil
			},
		},

		// Hex to other systems
		{
			Name:  "sent",
			Usage: "Convert a hex value to hexadecimal, binary, and octal",
			Action: func(c *cli.Context) error {
				var err error
				//var value string
				oflog.Init()
				if c.Args().Len() != 1 {
					oflog.Print.Fatalf("Please input a user id.")
					return nil
				}
				user := c.Args().Get(0)
				err = algorithm.IworkSent(user)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.IworkSent!")
					return err
				}
				return nil
			},
		},
	},
}
