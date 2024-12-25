package cmd

import (
	"fmt"
	"oftools/algorithm"

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
			Usage: "Convert a hex value to hexadecimal, binary, and octal",
			Action: func(ctx *cli.Context) error {
				value := ctx.Args().Get(0)
				var result algorithm.Conversion
				var err error
				result, err = algorithm.ConvertHexToOthers(value)
				if err != nil {
					return fmt.Errorf("invalid hexadecimal value: %v", err)
				}
				fmt.Printf("Hex: %s -> Hex: 0x%X\n", value, result.Hex)
				fmt.Printf("Hex: %s -> Dec: %d\n", value, result.Dec)
				fmt.Printf("Hex: %s -> Oct: %s\n", value, result.Oct)
				fmt.Printf("Hex: %s -> Bin: %s\n", value, result.Bin)
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
				result, err = algorithm.ConvertHexToOthers(value)
				if err != nil {
					return fmt.Errorf("invalid decimal value: %v", err)
				}
				fmt.Printf("Dec: %s -> Hex: 0x%X\n", value, result.Hex)
				fmt.Printf("Dec: %s -> Dec: %d\n", value, result.Dec)
				fmt.Printf("Dec: %s -> Oct: %s\n", value, result.Oct)
				fmt.Printf("Dec: %s -> Bin: %s\n", value, result.Bin)
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
				result, err = algorithm.ConvertHexToOthers(value)
				if err != nil {
					return fmt.Errorf("invalid octal value: %v", err)
				}
				fmt.Printf("Oct: %s -> Hex: 0x%X\n", value, result.Hex)
				fmt.Printf("Oct: %s -> Dec: %d\n", value, result.Dec)
				fmt.Printf("Oct: %s -> Oct: %s\n", value, result.Oct)
				fmt.Printf("Oct: %s -> Bin: %s\n", value, result.Bin)
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
				result, err = algorithm.ConvertHexToOthers(value)
				if err != nil {
					return fmt.Errorf("invalid binary value: %v", err)
				}
				fmt.Printf("Bin: %s -> Hex: 0x%X\n", value, result.Hex)
				fmt.Printf("Bin: %s -> Dec: %d\n", value, result.Dec)
				fmt.Printf("Bin: %s -> Oct: %s\n", value, result.Oct)
				fmt.Printf("Bin: %s -> Bin: %s\n", value, result.Bin)
				return nil
			},
		},
	},
}
