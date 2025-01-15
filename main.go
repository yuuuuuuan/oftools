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
		HideHelpCommand:      true,
		Commands: []*cli.Command{
			cmd.GreetCommand,   // Register greet command
			cmd.CalcCommand,    // Register calc command
			cmd.VersionCommand, // Register version command
			cmd.ConvertCommand,
			cmd.ExcelCommand,
			cmd.JumpCommand,
			cmd.ProgramCommand,
			cmd.HttpCommand,
			cmd.SignCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		oflog.Print.Fatal(err)
	}
	// oflog.Print.Infof("123")
	// oflog.Print.Errorf("123")
	// oflog.Print.Fatalf("123")
	// err1 := error1()
	// if err1 != nil {
	// 	oflog.Print.Errorf("%s:something wrong at %v", getFunctionName(), err1)
	// }
	// oflog.Print.Errorf("Error")
	// oflog.Print.Infof("Error")
}

// func error1() error {
// 	oflog.Init()
// 	err := errors.New("error:err1")
// 	oflog.Print.Errorf("%s:something wrong:%v", getFunctionName(), err)
// 	return err
// }

// // GetFunctionName retrieves the name of the currently executing function
// func getFunctionName() string {
// 	pc, _, _, _ := runtime.Caller(1)
// 	funcObj := runtime.FuncForPC(pc)
// 	return funcObj.Name()
// }
//GOOS=windows GOARCH=amd64 go build -o oftools.exe
//go build main.go
