package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"
	"fmt"
	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var OaCommand = &cli.Command{
	Name:            "oa",
	Aliases:         []string{"o"},
	Usage:           "🦋 Ofilm oa system bash shell",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "results",
			Usage: "🍿 Get a user results from a year ago to now.",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()

				if ctx.Args().Len() == 1 {
					user := ctx.Args().Get(0)
					err = algorithm.OaResults(user)
					if err != nil {
						fmt.Println(err)
						oflog.Print.Fatalf("Function start failed at algorithm.OaResults!")
						return err
					}
				} else {
					oflog.Print.Fatalf("Do not support one more arg.")
					return err
				}

				return nil
			},
		},

		{
			Name:  "info",
			Usage: "🏀 Get a user info from oa.",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()

				if ctx.Args().Len() == 1 {
					user := ctx.Args().Get(0)
					err = algorithm.OaInfo(user)
					if err != nil {
						fmt.Println(err)
						oflog.Print.Fatalf("Function start failed at algorithm.OaInfo!")
						return err
					}
				} else {
					oflog.Print.Fatalf("Do not support one more arg.")
					return err
				}

				return nil
			},
		},
	},
}
