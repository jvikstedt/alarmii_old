package action

import (
	"bufio"
	"bytes"
	"encoding/json"
	"testing"

	"github.com/jvikstedt/alarmii/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type JobServiceMock struct {
	mock.Mock
}

func (m *JobServiceMock) Add(job models.Job) (models.Job, error) {
	args := m.Called(job)
	return args.Get(0).(models.Job), args.Error(1)
}
func (m *JobServiceMock) Update(job models.Job) (models.Job, error) {
	args := m.Called(job)
	return args.Get(0).(models.Job), args.Error(1)
}
func (m *JobServiceMock) Find(id int) (models.Job, error) {
	args := m.Called(id)
	return args.Get(0).(models.Job), args.Error(1)
}
func (m *JobServiceMock) Delete(id int) (models.Job, error) {
	args := m.Called(id)
	return args.Get(0).(models.Job), args.Error(1)
}
func (m *JobServiceMock) All() ([]models.Job, error) {
	args := m.Called()
	return args.Get(0).([]models.Job), args.Error(1)
}

func TestList(t *testing.T) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	jobServiceMock := new(JobServiceMock)
	jobs := []models.Job{models.Job{Timing: "@every 15"}}
	jobServiceMock.On("All").Return(jobs, nil)
	job := NewJob(writer, jobServiceMock)
	job.List(Context{})
	writer.Flush()
	var resultJobs []models.Job
	json.Unmarshal(b.Bytes(), &resultJobs)
	assert.Equal(t, len(jobs), len(resultJobs))
	assert.Equal(t, jobs[0].Timing, resultJobs[0].Timing)
}
