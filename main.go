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
	config := LoadConfig("config.json")
	output, _ := exec.Command(config.Projects[0].Jobs[0].Command, config.Projects[0].Jobs[0].Arguments...).Output()
	log.Info(config)
	log.Info(string(output))
}
