package models

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/boltdb/bolt"
	"github.com/jvikstedt/alarmii/helper"
)

// Job struct that defines a job
type Job struct {
	ID             int               `json:"id"`
	ProjectID      int               `json:"project_id"`
	Timing         string            `json:"timing"`
	Command        string            `json:"command"`
	Arguments      []string          `json:"arguments"`
	ExpectedResult map[string]string `json:"expected_result"`
}

var jobsBucket = []byte("jobs")

// SaveJob saves a job to database
func (j *Job) SaveJob() (err error) {
	err = Database.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(jobsBucket)
		if err != nil {
			return err
		}
		if j.ID == 0 {
			id, err := b.NextSequence()
			if err != nil {
				return err
			}
			j.ID = int(id)
		}

		encoded, err := json.Marshal(j)
		if err != nil {
			return err
		}
		return b.Put(helper.Itob(j.ID), encoded)
	})
	return
}

// GetJobs gets all jobs from database
func GetJobs() (jobs []Job, err error) {
	err = Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(jobsBucket)
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			var job Job
			err = json.Unmarshal(v, &job)
			if err != nil {
				return err
			}
			jobs = append(jobs, job)
			return nil
		})
		return err
	})
	return
}

// GetJobByID gets single job id from database
func GetJobByID(id int) (job Job, err error) {
	err = Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(jobsBucket)
		if b == nil {
			return nil
		}
		bytes := b.Get(helper.Itob(id))
		err = json.Unmarshal(bytes, &job)
		return err
	})
	return
}

// DeleteJobByID deletes job by id
func DeleteJobByID(id int) (err error) {
	err = Database.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(jobsBucket)
		if err != nil {
			return err
		}
		return b.Delete(helper.Itob(id))
	})
	return
}

// Runnable creates a Runnable that contains information about job
func (j Job) Runnable() Runnable {
	return Runnable{j.ID}
}

// Runnable struct that implements scheduler.Schedulable interface
// Reason using this instead of job itself is that this way we force
// database reload before running job and job scheduler does not keep all
// jobs in memory
type Runnable struct {
	JobID int
}

// Run handles loadin job by id from database and executing it
func (r Runnable) Run() {
	job, err := GetJobByID(r.JobID)
	if err != nil {
		return
	}

	output, _ := exec.Command(job.Command, job.Arguments...).Output()
	var objmap map[string]string
	json.Unmarshal(output, &objmap)
	for k, v := range job.ExpectedResult {
		log.Println(objmap[k] + " == " + v)
	}
}

// CronTime returns Timing for a job
func (r Runnable) CronTime() string {
	job, err := GetJobByID(r.JobID)
	if err != nil {
		return ""
	}
	return job.Timing
}
