package main

import (
	"encoding/json"
	"io/ioutil"
)

// Project struct that defines a project
type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Jobs        []Job  `json:"jobs"`
}

// Job struct that defines a job
type Job struct {
	Command        string            `json:"command"`
	Arguments      []string          `json:"arguments"`
	ExpectedResult map[string]string `json:"expected_result"`
}

// LoadProjects loads projects from specified file
// Expects json file
func LoadProjects(filePath string) (projects []Project) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &projects)
	return
}
