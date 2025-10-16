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
	Usage:           "🍀 Auto tools at wechat iwork",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "get",
			Usage: "🎾 Get a range of user name and id which are validate",
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
			Usage: "⚾ Sent a user name to get her/his private info",
			Action: func(c *cli.Context) error {
				var err error
				//var value string
				oflog.Init()
				if c.Args().Len() != 1 {
					oflog.Print.Fatalf("Please input a user id.")
					return err
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

		// Hex to other systems
		{
			Name:  "rencai",
			Usage: "🥎 Get the info of a user result from rencaipingding",
			Action: func(c *cli.Context) error {
				var err error
				//var value string
				oflog.Init()
				if c.Args().Len() != 1 {
					oflog.Print.Fatalf("Please input a user id.")
					return err
				}
				user := c.Args().Get(0)
				err = algorithm.IworkRencai(user)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.IworkRencai!")
					return err
				}
				return nil
			},
		},
	},
}
