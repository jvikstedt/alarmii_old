package models

import (
	"encoding/json"

	"github.com/boltdb/bolt"
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

// SaveProject saves a project to database
func SaveProject(project Project) (err error) {
	encoded, err := json.Marshal(project)
	if err != nil {
		return err
	}
	err = Database.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("projects"))
		if err != nil {
			return err
		}
		return b.Put([]byte(project.Name), encoded)
	})
	return
}

// GetProjects gets all projects from database
func GetProjects() (projects []Project, err error) {
	err = Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("projects"))
		b.ForEach(func(k, v []byte) error {
			var project Project
			err = json.Unmarshal(v, &project)
			if err != nil {
				return err
			}
			projects = append(projects, project)
			return nil
		})
		return err
	})
	return
}

// GetProjectByName gets single project by name from database
func GetProjectByName(name string) (project Project, err error) {
	err = Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("projects"))
		bytes := b.Get([]byte(name))
		err = json.Unmarshal(bytes, &project)
		return err
	})
	return
}
