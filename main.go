package main

import (
	"integration-cli/cmd"
	"log"
	"os"
)

func main() {
	app := cmd.InitCLI()
	cmd.Commands(app)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
