package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/jvikstedt/alarmii/commands"
	"github.com/jvikstedt/alarmii/helper"
	"github.com/jvikstedt/alarmii/models"
)

func init() {
	helper.SetupLogger("log/info.log")
	models.OpenDatabase("alarmii.db")
}

func main() {
	app := cli.NewApp()
	app.Name = "Alarmii"
	app.Usage = ""

	setupCommands(app)
	app.Run(os.Args)

	models.CloseDatabase()
}

func setupCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start alarmii process",
			Action:  commands.StartProcess,
		},
	}
}
