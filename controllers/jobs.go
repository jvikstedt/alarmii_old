package controllers

import (
	"net/http"
	"strconv"

	"github.com/jvikstedt/alarmii/models"
	"github.com/labstack/echo"
)

// GetJobs retrieves all jobs
func GetJobs(c echo.Context) error {
	jobs, err := models.GetJobs()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, jobs)
}

// GetJob retrieves a single job by id
func GetJob(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	job, err := models.GetJobByID(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, job)
}

// CreateJob saves a job
func CreateJob(c echo.Context) error {
	job := &models.Job{}
	if err := c.Bind(job); err != nil {
		return err
	}
	if err := job.SaveJob(); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, job)
}

// UpdateJob updates a single job
func UpdateJob(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	job, err := models.GetJobByID(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Bind(&job); err != nil {
		return err
	}
	if err := job.SaveJob(); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, job)
}

// DeleteJob deletes a single job
func DeleteJob(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteJobByID(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
