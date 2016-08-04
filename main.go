package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/exec"

	"gopkg.in/urfave/cli.v1"

	log "github.com/Sirupsen/logrus"
	"github.com/jvikstedt/alarmii/models"
	"github.com/rifflock/lfshook"
	"github.com/robfig/cron"
	"github.com/yosssi/ace"
)

func setupLogger() {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  "log/info.log",
		log.ErrorLevel: "log/error.log",
	}))
}

func setupScheduler() {
	c := cron.New()
	projects, _ := models.GetProjects()
	for _, p := range projects {
		for _, j := range p.Jobs {
			if j.Timing == "" {
				continue
			}
			c.AddFunc(j.Timing, func() {
				output, _ := exec.Command(j.Command, j.Arguments...).Output()
				var objmap map[string]string
				json.Unmarshal(output, &objmap)
				for k, v := range j.ExpectedResult {
					log.Info(objmap[k] + " == " + v)
				}
			})
		}
	}
	c.Start()
}

func init() {
	setupLogger()
	models.OpenDatabase("alarmii.db")
	job := models.Job{Timing: "@every 10s", Command: "echo", Arguments: []string{"{\"status\":\"200\"}"}, ExpectedResult: map[string]string{"status": "200"}}
	var jobs []models.Job
	jobs = append(jobs, job)
	project := models.Project{Name: "Something", Description: "jep", Jobs: jobs}
	models.SaveProject(project)
	setupScheduler()
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
