package service

import "github.com/jvikstedt/alarmii/models"

type Job struct{}

func (j Job) Add(job models.Job) (models.Job, error) {
	return models.Job{}, nil
}
func (j Job) Update(job models.Job) (models.Job, error) {
	return models.Job{}, nil
}
func (j Job) Find(id int) (models.Job, error) {
	return models.Job{}, nil
}
func (j Job) Delete(id int) (models.Job, error) {
	return models.Job{}, nil
}
func (j Job) All() ([]models.Job, error) {
	return []models.Job{}, nil
}
