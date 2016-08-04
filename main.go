package main

import (
	"net/http"
	"os"

	"gopkg.in/urfave/cli.v1"

	log "github.com/Sirupsen/logrus"
	"github.com/jvikstedt/alarmii/models"
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

	//config := LoadConfig("config.json")
	//output, _ := exec.Command(config.Projects[0].Jobs[0].Command, config.Projects[0].Jobs[0].Arguments...).Output()
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
