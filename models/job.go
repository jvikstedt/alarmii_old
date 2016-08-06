package models

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
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
func SaveJob(job Job) (err error) {
	encoded, err := json.Marshal(job)
	if err != nil {
		return err
	}
	err = Database.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(jobsBucket)
		if err != nil {
			return err
		}
		id, err := b.NextSequence()
		job.ID = int(id)

		if err != nil {
			return err
		}
		return b.Put(helper.Itob(job.ID), encoded)
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

// CronTime returns timing
func (j Job) CronTime() string {
	return j.Timing
}

// Execute runs job
func (j Job) Execute() {
	log.Info(j.Command)
}
