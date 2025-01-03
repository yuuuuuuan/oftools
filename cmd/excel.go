package cmd

import (
	"fmt"
	"oftools/algorithm"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var ExcelCommand = &cli.Command{
	Name:    "excel",
	Aliases: []string{"e"},
	Usage:   "Deal with Data",
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "sumsingle",
			Usage: "Summary a single file to destinationDir",
			Action: func(ctx *cli.Context) error {
				var err error
				sourceDir := ctx.Args().Get(0)
				destDir := "D:\\.oftools\\excel\\work"
				err = algorithm.ExcelSumSinger(sourceDir, destDir)
				if err != nil {
					return fmt.Errorf("invalid decimal value: %v", err)
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "summult",
			Usage: "Summary sorts of files to destinationDir",
			Action: func(ctx *cli.Context) error {
				var err error
				sourceDirs := ctx.Args().Slice()
				destDir := "D:\\.oftools\\excel\\work"
				err = algorithm.ExcelSumMult(sourceDirs, destDir)
				if err != nil {
					return fmt.Errorf("invalid decimal value: %v", err)
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "sumself",
			Usage: "Summary a file itself",
			Action: func(ctx *cli.Context) error {
				var err error
				sourceDir := ctx.Args().Get(0)
				err = algorithm.ExcelSumSelf(sourceDir)
				if err != nil {
					return fmt.Errorf("invalid decimal value: %v", err)
				}
				return nil
			},
		},
		// Hex to other systems
		{
			Name:  "clear",
			Usage: "Clear workdir and save files to savedir",
			Action: func(ctx *cli.Context) error {
				var err error
				sourceDir := "D:\\.oftools\\excel\\work"
				destDir := "D:\\.oftools\\excel\\save"
				err = algorithm.ExcelClear(sourceDir, destDir)
				if err != nil {
					return fmt.Errorf("invalid decimal value: %v", err)
				}
				return nil
			},
		},
	},
}
