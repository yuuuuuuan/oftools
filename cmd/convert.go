package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var ConvertCommand = &cli.Command{
	Name:    "convert",
	Aliases: []string{"c"},
	Usage:   "Convert hex to dec",
	Subcommands: []*cli.Command{
		// Hex to other systems
		{
			Name:  "hex",
			Usage: "Convert a hex value to decimal, binary, and octal",
			Action: func(ctx *cli.Context) error {
				oflog.Init()
				value := ctx.Args().Get(0)
				var result algorithm.Conversion
				var err error
				result, err = algorithm.ConvertHexToOthers(value)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ConvertHexToOthers!")
					return err
				}
				oflog.Print.Infof("Hex: %s -> Hex: %s", value, result.Hex)
				oflog.Print.Infof("Hex: %s -> Dec: %d", value, result.Dec)
				oflog.Print.Infof("Hex: %s -> Oct: %s", value, result.Oct)
				oflog.Print.Infof("Hex: %s -> Bin: %s", value, result.Bin)
				return nil
			},
		},
		// Decimal to other systems
		{
			Name:  "dec",
			Usage: "Convert a decimal value to hex, binary, and octal",
			Action: func(ctx *cli.Context) error {
				value := ctx.Args().Get(0)
				var result algorithm.Conversion
				var err error
				result, err = algorithm.ConvertDecToOthers(value)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ConvertHexToOthers!")
					return err
				}
				oflog.Print.Infof("Hex: %s -> Hex: %s", value, result.Hex)
				oflog.Print.Infof("Hex: %s -> Dec: %d", value, result.Dec)
				oflog.Print.Infof("Hex: %s -> Oct: %s", value, result.Oct)
				oflog.Print.Infof("Hex: %s -> Bin: %s", value, result.Bin)
				return nil
			},
		},
		// Octal to other systems
		{
			Name:  "oct",
			Usage: "Convert a octal value to hex, binary, and decimal",
			Action: func(ctx *cli.Context) error {
				value := ctx.Args().Get(0)
				var result algorithm.Conversion
				var err error
				result, err = algorithm.ConvertOctToOthers(value)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ConvertOctToOthers!")
					return err
				}
				oflog.Print.Infof("Hex: %s -> Hex: %s", value, result.Hex)
				oflog.Print.Infof("Hex: %s -> Dec: %d", value, result.Dec)
				oflog.Print.Infof("Hex: %s -> Oct: %s", value, result.Oct)
				oflog.Print.Infof("Hex: %s -> Bin: %s", value, result.Bin)
				return nil
			},
		},
		// Binary to other systems
		{
			Name:  "bin",
			Usage: "Convert a binary value to hex, octal, and decimal",
			Action: func(ctx *cli.Context) error {
				value := ctx.Args().Get(0)
				var result algorithm.Conversion
				var err error
				result, err = algorithm.ConvertBinToOthers(value)
				if err != nil {
					oflog.Print.Fatalf("Function start failed at algorithm.ConvertBinToOthers!")
					return err
				}
				oflog.Print.Infof("Hex: %s -> Hex: %s", value, result.Hex)
				oflog.Print.Infof("Hex: %s -> Dec: %d", value, result.Dec)
				oflog.Print.Infof("Hex: %s -> Oct: %s", value, result.Oct)
				oflog.Print.Infof("Hex: %s -> Bin: %s", value, result.Bin)
				return nil
			},
		},
	},
}
