package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveJob(t *testing.T) {
	job := Job{Timing: "@every 10s"}
	err := job.SaveJob()
	assert.Nil(t, err)
}

func TestGetJobs(t *testing.T) {
	jobs, err := GetJobs()
	assert.Nil(t, err)
	assert.Equal(t, len(jobs), 1)
	assert.Equal(t, jobs[0].Timing, "@every 10s")
}

func TestGetJobsByID(t *testing.T) {
	job, err := GetJobByID(1)
	assert.Nil(t, err)
	assert.Equal(t, job.Timing, "@every 10s")
}

func TestDeleteJobByID(t *testing.T) {
	err := DeleteJobByID(1)
	assert.Nil(t, err)

	jobs, err := GetJobs()
	assert.Nil(t, err)
	assert.Equal(t, len(jobs), 0)
}
