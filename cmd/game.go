package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var GameCommand = &cli.Command{
	Name:            "game",
	Aliases:         []string{"gm"},
	Usage:           "🎮 Some intersting game.",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "wordle",
			Usage: "🎲 Bulls and Cows Game.",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()

				if ctx.Args().Len() == 0 {
					err = algorithm.GameWordle()
					if err != nil {
						oflog.Print.Fatalf("Function start failed at algorithm.GameWordle!")
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
