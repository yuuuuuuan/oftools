package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var GreetCommand = &cli.Command{
	Name:    "greet",
	Aliases: []string{"g"},
	Usage:   "🌙 Greets the user with the provided name",
	HideHelpCommand:      true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Name of the person to greet",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		name := c.String("name")
		fmt.Printf("Hello, %s!\n", name)
		return nil
	},
}
