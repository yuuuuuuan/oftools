package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

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
			Usage: "🍿 Get quiz with token and Testpaper id.",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()

				if ctx.Args().Len() == 1 {
					err = algorithm.OaResults()
					if err != nil {
						println(err)
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
	},
}
