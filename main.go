package main

import (
	"net/http"
	"os"

	"gopkg.in/urfave/cli.v1"

	log "github.com/Sirupsen/logrus"
	"github.com/jvikstedt/alarmii/models"
	"github.com/jvikstedt/alarmii/scheduler"
	"github.com/rifflock/lfshook"
	"github.com/yosssi/ace"
)

func setupLogger() {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  "log/info.log",
		log.ErrorLevel: "log/error.log",
	}))
}

func init() {
	setupLogger()
	models.OpenDatabase("alarmii.db")
	job := models.Job{Timing: "@every 10s", Command: "echo", Arguments: []string{"{\"status\":\"200\"}"}, ExpectedResult: map[string]string{"status": "200"}}
	models.SaveJob(&job)

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
