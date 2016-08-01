package main

import (
	"encoding/json"
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
	testJSON := []byte(`{
		"name": "TestProject",
		"description": "Testing...",
		"jobs": [
		{"command": "echo", "arguments": ["-n", "{\"status\":\"200\"}"], "expected_result": {"status": "200"}}
		]
	}`)
	var project Project
	json.Unmarshal(testJSON, &project)
	log.Info(project)
	output, _ := exec.Command(project.Jobs[0].Command, project.Jobs[0].Arguments...).Output()
	log.Info(string(output))
}
