package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Remove("tmp/test.db")
	OpenDatabase("/tmp/test.db")
	retCode := m.Run()
	CloseDatabase()
	os.Exit(retCode)
}

func TestSaveProject(t *testing.T) {
	project := Project{Name: "Something", Description: "Cool"}
	err := SaveProject(project)
	assert.Nil(t, err)
}

func TestGetProjects(t *testing.T) {
	projects, err := GetProjects()
	assert.Nil(t, err)
	assert.Equal(t, len(projects), 1)
	assert.Equal(t, projects[0].Name, "Something")
}

func TestGetProjectByName(t *testing.T) {
	project, err := GetProjectByName("Something")
	assert.Nil(t, err)
	assert.Equal(t, project.Name, "Something")
	assert.Equal(t, project.Description, "Cool")
}
