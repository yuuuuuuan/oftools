package main

import (
	"oftools/cmd"
	"oftools/oflog"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	oflog.Init()
	app := &cli.App{
		Name:                 "oftools",
		Usage:                "fight the loneliness!",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			cmd.GreetCommand,   // Register greet command
			cmd.CalcCommand,    // Register calc command
			cmd.VersionCommand, // Register version command
			cmd.ConvertCommand,
			cmd.ExcelCommand,
			cmd.JumpCommand,
			cmd.ProgramCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		oflog.Print.Fatal(err)
	}
	// oflog.Print.Errorf("Error")
	// oflog.Print.Infof("Error")
}
