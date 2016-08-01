package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config config contains configuration options
type Config struct {
	Projects []Project `json:"projects"`
}

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

// LoadConfig loads config from specified file
// Expects json file
func LoadConfig(filePath string) (config Config) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &config)
	return
}
