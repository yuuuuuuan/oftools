package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// 🎉🌟🐱🚀🍕🍀💖🦄🌸💫🌙🌿🍓⚡️🎶🍩🎮🎯🎤💎🍒🏞️🎀🛸🎠👑🐾🌈🔥🧩🍔🍉🎨🌻🎡🍭✨🍄🎈📚🌍🎧🍪🍊🍍🍎🍑🍒🍓🍉🍇🍈🍋🍌🍍🍕🌮🥗🍜🍱🍘🍫🍬🍪🍩🥞🍦🍧🍨🍮🍿🎯🧸🛸🚲🚗🛳️🏖️⛷️⚽
// GreetCommand is a simple command to greet the user.
var ExcelCommand = &cli.Command{
	Name:    "excel",
	Aliases: []string{"e"},
	Usage:   "🎉 Deal with Data",
	HideHelpCommand:      true,
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "sumsingle",
			Usage: "🌸 Summary a single file to destinationDir",
			Action: func(ctx *cli.Context) error {
				var err error
				oflog.Init()
				if ctx.Args().Len() != 1 {
					oflog.Print.Fatalf("Please input a single file Dir!")
					return nil
				}
				sourceDir := ctx.Args().Get(0)
				destDir := "D:\\.oftools\\excel\\work"
				err = algorithm.ExcelSumSinger(sourceDir, destDir)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ExcelSumSinger!")
					return err
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "summult",
			Usage: "💫 Summary sorts of files to destinationDir",
			Action: func(ctx *cli.Context) error {
				var err error
				if ctx.Args().Len() == 0 {
					oflog.Print.Fatalf("Please input one or sorts of files Dir!")
					return nil
				}
				sourceDirs := ctx.Args().Slice()
				destDir := "D:\\.oftools\\excel\\work"
				err = algorithm.ExcelSumMult(sourceDirs, destDir)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ExcelSumMult!")
					return err
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "sumself",
			Usage: "🌿 Summary a file itself",
			Action: func(ctx *cli.Context) error {
				var err error
				if ctx.Args().Len() != 1 {
					oflog.Print.Fatalf("Please input a single file Dir!")
					return nil
				}
				sourceDir := ctx.Args().Get(0)
				err = algorithm.ExcelSumSelf(sourceDir)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ExcelSumSelf!")
					return err
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "delfail",
			Usage: "🎠 Delete repfail and repeat Data",
			Action: func(ctx *cli.Context) error {
				var err error
				sourceDir := "D:\\.oftools\\excel\\work"
				destDir := "D:\\.oftools\\excel\\save"
				err = algorithm.ExcelClear(sourceDir, destDir)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ExcelClear!")
					return err
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "sumcross",
			Usage: "🍔 Summary different Data by model ID",
			Action: func(ctx *cli.Context) error {
				var err error
				sourceDir := "D:\\.oftools\\excel\\work"
				destDir := "D:\\.oftools\\excel\\save"
				err = algorithm.ExcelClear(sourceDir, destDir)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ExcelClear!")
					return err
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "clear",
			Usage: "🍓 Clear workdir and save files to savedir",
			Action: func(ctx *cli.Context) error {
				var err error
				sourceDir := "D:\\.oftools\\excel\\work"
				destDir := "D:\\.oftools\\excel\\save"
				err = algorithm.ExcelClear(sourceDir, destDir)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ExcelClear!")
					return err
				}
				return nil
			},
		},
	},
}
