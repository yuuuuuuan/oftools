package cmd

import (
	"oftools/algorithm"
	"oftools/oflog"

	"github.com/urfave/cli/v2"
)

// GreetCommand is a simple command to greet the user.
var HttpCommand = &cli.Command{
	Name:    "http",
	Aliases: []string{"ht"},
	Usage:   "Greets the user with the provided name",
	Action: func(c *cli.Context) error {
		//var err error
		oflog.Init()
		port := c.Args().Get(0)
		err := algorithm.Http(port)
		if err != nil {
			oflog.Print.Fatalf("Function start failed at algorithm.Http!")
			return err
		}
		return nil
	},
}
