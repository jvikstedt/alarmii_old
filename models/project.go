package models

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/jvikstedt/alarmii/helper"
)

// Project struct that defines a project
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var projectsBucket = []byte("projects")

// SaveProject saves a project to database
func (p *Project) SaveProject() (err error) {
	err = Database.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(projectsBucket)
		if err != nil {
			return err
		}
		if p.ID == 0 {
			id, err := b.NextSequence()
			if err != nil {
				return err
			}
			p.ID = int(id)
		}

		encoded, err := json.Marshal(p)
		if err != nil {
			return err
		}

		return b.Put(helper.Itob(p.ID), encoded)
	})
	return
}

// GetProjects gets all projects from database
func GetProjects() (projects []Project, err error) {
	err = Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(projectsBucket)
		if b == nil {
			return nil
		}
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

// GetProjectByID gets single project id from database
func GetProjectByID(id int) (project Project, err error) {
	err = Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(projectsBucket)
		if b == nil {
			return nil
		}
		bytes := b.Get(helper.Itob(id))
		err = json.Unmarshal(bytes, &project)
		return err
	})
	return
}

// DeleteProjectByID deletes project by id
func DeleteProjectByID(id int) (err error) {
	err = Database.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(projectsBucket)
		if err != nil {
			return err
		}
		return b.Delete(helper.Itob(id))
	})
	return
}
