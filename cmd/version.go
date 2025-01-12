package cmd

import (
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// VersionCommand is a command that shows the version of the app.
var VersionCommand = &cli.Command{
	Name:            "version",
	Aliases:         []string{"v"},
	Usage:           "🌸 Prints the version of the app",
	HideHelpCommand: true,
	Action: func(c *cli.Context) error {
		oflog.Init()
		oflog.Print.Warnf("oftools version 1.0.0")
		return nil
	},
}
