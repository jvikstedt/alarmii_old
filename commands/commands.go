package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/jvikstedt/alarmii/controllers"
	"github.com/jvikstedt/alarmii/helper"
	"github.com/jvikstedt/alarmii/models"
	"github.com/jvikstedt/alarmii/scheduler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/tylerb/graceful"
	cli "gopkg.in/urfave/cli.v1"
)

// StartProcess starts persistent process
func StartProcess(c *cli.Context) (err error) {
	models.OpenDatabase("alarmii.db")
	defer models.CloseDatabase()

	helper.SavePID()
	scheduler.SetupScheduler()
	jobs, err := models.GetJobs()
	if err != nil {
		return
	}
	fmt.Println("Starting following jobs:")
	for _, v := range jobs {
		scheduler.AddSchedulable(v.Runnable())
		fmt.Println(v)
	}
	scheduler.StartScheduler()
	defer scheduler.StopScheduler()

	runServer()

	return
}

// StopProcess starts persistent process
func StopProcess(c *cli.Context) (err error) {
	pid, err := helper.ReadPID()
	if err != nil {
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return
	}
	err = process.Signal(os.Interrupt)
	return
}

func runServer() {
	e := echo.New()
	g := e.Group("/api/v1")
	g.GET("/jobs", controllers.GetJobs)
	g.GET("/jobs/:id", controllers.GetJob)
	g.POST("/jobs", controllers.CreateJob)
	g.PATCH("/jobs/:id", controllers.UpdateJob)
	g.DELETE("/jobs/:id", controllers.DeleteJob)
	std := standard.New(":3000")
	std.SetHandler(e)
	graceful.ListenAndServe(std.Server, 5*time.Second)
}

// ListJobs list all jobs from database
func ListJobs(c *cli.Context) (err error) {
	jobs, err := models.GetJobs()
	if err != nil {
		return
	}
	for _, v := range jobs {
		fmt.Println(v)
	}
	return
}
