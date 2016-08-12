package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/jvikstedt/alarmii/commands"
	"github.com/jvikstedt/alarmii/helper"
)

func init() {
	helper.SetupLogger("log/info.log")
}

func main() {
	app := cli.NewApp()
	app.Name = "Alarmii"
	app.Usage = ""

	setupCommands(app)
	app.Run(os.Args)
}

func setupCommands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start alarmii process",
			Action:  commands.StartProcess,
		},
		{
			Name:   "stop",
			Usage:  "Stop alarmii process",
			Action: commands.StopProcess,
		},
		{
			Name: "job",
			Subcommands: []cli.Command{
				{
					Name:   "list",
					Action: commands.ListJobs,
				},
			},
		},
	}
}
