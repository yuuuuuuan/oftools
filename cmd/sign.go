package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var SignCommand = &cli.Command{
	Name:            "sign",
	Aliases:         []string{"si"},
	Usage:           "🍭 Deal with Program",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "pingpong",
			Usage: "✨ Sign at Pingpong activity",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()
				//sourceDir := ctx.Args().Get(0)
				//firewareDir := "D:\\.oftools\\excel\\work"
				err = algorithm.SignPingpong()
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.SignPingpong!")
					return err
				}
				return nil
			},
		},
		{
			Name:  "badminton",
			Usage: "✨ Sign at Badminton activity",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()
				//sourceDir := ctx.Args().Get(0)
				//firewareDir := "D:\\.oftools\\excel\\work"
				err = algorithm.SignBadminton()
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.SignBadminton!")
					return err
				}
				return nil
			},
		},
		{
			Name:  "test",
			Usage: "🍄 Test for Signnig at activity",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()
				//sourceDir := ctx.Args().Get(0)
				//firewareDir := "D:\\.oftools\\excel\\work"
				err = algorithm.SignTest()
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.SignTest!")
					return err
				}
				return nil
			},
		},
	},
}
