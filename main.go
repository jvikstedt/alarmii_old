package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/jvikstedt/alarmii/helper"
	"github.com/jvikstedt/alarmii/models"
	"github.com/jvikstedt/alarmii/scheduler"
	"github.com/yosssi/ace"
)

func init() {
	helper.SetupLogger("log/info.log")
	models.OpenDatabase("alarmii.db")
	job := models.Job{Timing: "@every 10s", Command: "echo", Arguments: []string{"{\"status\":\"200\"}"}, ExpectedResult: map[string]string{"status": "200"}}
	job.SaveJob()

	scheduler.SetupScheduler()
	scheduler.AddSchedulable(job.Runnable())
	scheduler.StartScheduler()
}

func main() {
	app := cli.NewApp()
	app.Name = "Alarmii"
	app.Usage = ""
	app.Action = func(c *cli.Context) error {
		startServer()
		return nil
	}
	app.Run(os.Args)

	models.CloseDatabase()
}

func test(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("templates/layout", "templates/content", nil)
	projects, _ := models.GetProjects()
	data := map[string]interface{}{
		"Projects": projects,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startServer() {
	http.HandleFunc("/", test)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
