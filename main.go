package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/jvikstedt/alarmii/helper"
	"github.com/jvikstedt/alarmii/models"
	"github.com/jvikstedt/alarmii/scheduler"
)

var running = true

func main() {
	app := cli.NewApp()
	app.Name = "Alarmii"
	app.Usage = ""
	app.Action = func(c *cli.Context) error {
		helper.SetupLogger("log/info.log")
		models.OpenDatabase("alarmii.db")

		job := models.Job{Timing: "@every 10s", Command: "echo", Arguments: []string{"{\"status\":\"200\"}"}, ExpectedResult: map[string]string{"status": "200"}}
		job.SaveJob()

		scheduler.SetupScheduler()
		scheduler.AddSchedulable(job.Runnable())
		scheduler.StartScheduler()
		for running {
		}
		return nil
	}
	app.Run(os.Args)

	scheduler.StopScheduler()
	models.CloseDatabase()
}
