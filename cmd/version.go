package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// VersionCommand is a command that shows the version of the app.
var VersionCommand = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Prints the version of the app",
	Action: func(c *cli.Context) error {
		fmt.Println("oftools version 1.0.0")
		return nil
	},
}
