package main

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
