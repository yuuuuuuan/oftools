package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var OfyxCommand = &cli.Command{
	Name:            "ofyx",
	Aliases:         []string{"of"},
	Usage:           "🖼️  Ofilm yun xue bash shell",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "getquiz",
			Usage: "🏈 Get quiz with token and Testpaper id.",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()

				if ctx.Args().Len() == 0 {
					err = algorithm.OfyxGetquiz()
					if err != nil {
						oflog.Print.Fatalf("Function start failed at algorithm.OfyxGetquiz!")
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
