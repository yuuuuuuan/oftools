package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var OutlookCommand = &cli.Command{
	Name:            "outlook",
	Aliases:         []string{"o"},
	Usage:           "👑 Greets the user with the provided name",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "listen",
			Usage: "Listen outlook email",
			Action: func(c *cli.Context) error {
				var err error
				//var value string
				oflog.Init()
				err = algorithm.OutlookListen()
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.IworkGet!")
					return err
				}
				return nil
			},
		},
	},
}
