package main

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

func setupLogger() {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  "log/info.log",
		log.ErrorLevel: "log/error.log",
	}))
}

func init() {
	setupLogger()
}

func main() {
	projects := LoadProjects("test.json")
	output, _ := exec.Command(projects[0].Jobs[0].Command, projects[0].Jobs[0].Arguments...).Output()
	log.Info(projects)
	log.Info(string(output))
}
