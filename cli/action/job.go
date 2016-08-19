package action

import (
	"encoding/json"
	"io"

	"github.com/jvikstedt/alarmii/models"
)

// JobService defines required functions
type JobService interface {
	Add(job models.Job) (models.Job, error)
	Update(job models.Job) (models.Job, error)
	Find(id int) (models.Job, error)
	Delete(id int) (models.Job, error)
	All() ([]models.Job, error)
}

// Job is a Action for Jobs
type Job struct {
	Action
	Service JobService
}

// NewJob constructs a Job
func NewJob(output io.Writer, jobService JobService) Job {
	return Job{Action: Action{Output: output}, Service: jobService}
}

// Add creates a new job to the output
func (j Job) Add(c Context) error {
	j.WriteString("Adding...")
	return nil
}

// List lists all jobs to the output
func (j Job) List(c Context) error {
	jobs, err := j.Service.All()
	if err != nil {
		return err
	}
	jobsPretty, _ := json.MarshalIndent(jobs, "", "  ")
	j.Output.Write(jobsPretty)
	return nil
}
