package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// GreetCommand is a simple command to greet the user.
var ConvertCommand = &cli.Command{
	Name:    "Convert",
	Aliases: []string{"c"},
	Usage:   "Greets the user with the provided name",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Name of the person to greet",
			Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		name := c.String("name")
		fmt.Printf("Hello, %s!\n", name)
		return nil
	},
}
