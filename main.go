package main

import (
	"log"

	"oftools/cmd"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "oftools",
		Usage: "fight the loneliness!",
		Commands: []*cli.Command{
			cmd.GreetCommand,   // Register greet command
			cmd.CalcCommand,    // Register calc command
			cmd.VersionCommand, // Register version command
			cmd.ConvertCommand,
			cmd.ExcelCommand,
			cmd.JumpCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
