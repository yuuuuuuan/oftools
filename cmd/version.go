package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// VersionCommand is a command that shows the version of the app.
var VersionCommand = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Prints the version of the app",
	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Println("oftools version 1.0.0")
		return nil
	},
}
