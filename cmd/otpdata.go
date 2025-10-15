package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var OtpDataCommand = &cli.Command{
	Name:            "otpdata",
	Aliases:         []string{"otp"},
	Usage:           "🍟 Sorting Otp Data",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "single",
			Usage: "Sorting a single otp ini file",
			Action: func(ctx *cli.Context) error {
				var err error
				//var value string
				oflog.Init()

				if ctx.Args().Len() != 1 {
					oflog.Print.Fatalf("Please input one or sorts of files Dir!")
					return nil
				}
				filepath := ctx.Args().Get(0)

				nums := []string{"9", "9819", "9820", "9821", "9822", "9823", "9824", "9838", "9840", "9891", "9892", "9893", "9894", "9895", "9896", "9913", "9914", "9915", "9916", "9917", "9918"}
				err = algorithm.OtpdataGetSingle(filepath, nums)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.OtpdataGetSingle!")
					return err
				}
				return nil
			},
		},

		{
			Name:  "muti",
			Usage: "Sorting series of single otp ini file",
			Action: func(ctx *cli.Context) error {
				var err error
				//var value string
				oflog.Init()

				if ctx.Args().Len() != 1 {
					oflog.Print.Fatalf("Please input one or sorts of files Dir!")
					return nil
				}
				filepath := ctx.Args().Get(0)

				err = algorithm.OtpdataGetMuti(filepath)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.OtpdataGetMuti!")
					return err
				}
				return nil
			},
		},
	},
}
