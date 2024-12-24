package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// CalcCommand is a simple calculator command.
var CalcCommand = &cli.Command{
	Name:    "calc",
	Aliases: []string{"cal"},
	Usage:   "Performs basic arithmetic operations (add, subtract, multiply, divide)",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "operation",
			Aliases:  []string{"o"},
			Usage:    "Operation to perform (add, sub, mul, div)",
			Required: false,
		},
		&cli.Float64Flag{
			Name:     "a",
			Aliases:  []string{"x"},
			Usage:    "First number",
			Required: true,
		},
		&cli.Float64Flag{
			Name:     "b",
			Aliases:  []string{"y"},
			Usage:    "Second number",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		a := c.Float64("a")
		b := c.Float64("b")
		operation := c.String("operation")

		var result float64
		var err error

		switch operation {
		case "add":
			result = a + b
		case "sub":
			result = a - b
		case "mul":
			result = a * b
		case "div":
			if b == 0 {
				err = fmt.Errorf("cannot divide by zero")
			} else {
				result = a / b
			}
		default:
			err = fmt.Errorf("invalid operation: %s", operation)
		}

		if err != nil {
			return err
		}

		fmt.Printf("Result: %f\n", result)
		return nil
	},
}
