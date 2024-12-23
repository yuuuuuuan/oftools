package main

import (
    "fmt"
    "log"
    "os"
    "context"

    "github.com/urfave/cli/v3"
)

func main() {
    app := &cli.Command{
        Name:  "oftools",
        Usage: "fight the loneliness!",
        Action: func(context.Context, *cli.Command) error {
            fmt.Println("Hello friend!")
            return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}