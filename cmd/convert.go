package cmd

import (
	"fmt"
	"strconv"
	"strings"
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
			Usage: "Convert a hex value to decimal, binary, and octal",
			Action: func(ctx *cli.Context) error {
				hexValue := ctx.Args().Get(0)

				algorithm.ConvertHexToOthers(hexValue)

				return nil
			},
		},
		// Decimal to other systems
		{
			Name:  "decimal",
			Usage: "Convert a decimal value to hex, binary, and octal",
			Action: func(ctx *cli.Context) error {
				decimalValueStr := ctx.Args().Get(0)
				decimalValue, err := strconv.ParseInt(decimalValueStr, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid decimal value: %v", err)
				}

				// Convert and print to different systems
				fmt.Printf("Decimal: %d -> Hex: 0x%X\n", decimalValue, decimalValue)
				fmt.Printf("Decimal: %d -> Binary: %b\n", decimalValue, decimalValue)
				fmt.Printf("Decimal: %d -> Octal: %o\n", decimalValue, decimalValue)

				return nil
			},
		},
	},
}
