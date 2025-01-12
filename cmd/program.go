package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var ProgramCommand = &cli.Command{
	Name:    "program",
	Aliases: []string{"p"},
	Usage:   "🦄 Deal with Program",
	HideHelpCommand:      true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "firmwaresingle",
			Usage: "change the OIS firmware",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()
				sourceDir := ctx.Args().Get(0)
				//firewareDir := "D:\\.oftools\\excel\\work"
				err = algorithm.ProgramFirewareSingle(sourceDir)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ProgramFirewareSingle!")
					return err
				}
				return nil
			},
		},
	},
}
