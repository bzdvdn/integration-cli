package cmd

import (
	"integration-cli/structure"
	"log"
	"os/exec"

	"github.com/urfave/cli"
)

// InitCLI ...
func InitCLI() *cli.App {
	return cli.NewApp()
}

// Commands ...
func Commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Init integraton",
			Action: func(c *cli.Context) {
				var name string
				if c.NArg() == 0 {
					log.Fatal("Enter name for your project")
				}
				name = c.Args()[0]
				structure.GenerateStucts(name)
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build go app",
			Action: func(c *cli.Context) {
				cmd := exec.Command("make", "build")
				log.Printf("Building ...")
				err := cmd.Run()
				log.Printf("Build finished with error: %v", err)
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run go app",
			Action: func(c *cli.Context) {
				cmd := exec.Command("make", "run")
				log.Printf("Runnin ...")
				err := cmd.Run()
				log.Printf("Finis running with error: %v", err)
			},
		},
	}
}
