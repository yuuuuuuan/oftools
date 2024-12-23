package main

import (
	"context"
	"log"
	"oftools/cmd"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "oftools",
		Usage: "fight the loneliness!",
		Commands: []*cli.Command{
			cmd.GreetCommand,   // Register greet command
			cmd.CalcCommand,    // Register calc command
			cmd.VersionCommand, // Register version command
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
