package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jvikstedt/alarmii/controllers"
	"github.com/jvikstedt/alarmii/helper"
	"github.com/jvikstedt/alarmii/models"
	"github.com/jvikstedt/alarmii/scheduler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/parnurzeal/gorequest"
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

// ListJobs list all jobs from the server
func ListJobs(c *cli.Context) error {
	request, body, errs := gorequest.New().Get("http://localhost:3000/api/v1/jobs").End()
	if len(errs) != 0 {
		return cli.NewExitError(errs[0].Error(), 1)
	}
	if request.StatusCode != 200 {
		return cli.NewExitError(fmt.Sprintf("Something went wrong, status: %d", request.StatusCode), 1)
	}
	var jobs []models.Job
	json.Unmarshal([]byte(body), &jobs)
	jobsPretty, _ := json.MarshalIndent(jobs, "", "  ")
	fmt.Println(string(jobsPretty))
	return nil
}

// AddJob sends job to server
func AddJob(c *cli.Context) error {
	var job models.Job
	if timing := c.String("timing"); timing != "" {
		job.Timing = timing
	}
	if command := c.String("command"); command != "" {
		job.Command = command
	}
	if arguments := c.StringSlice("arguments"); len(arguments) > 0 {
		job.Arguments = arguments
	}
	if er := c.String("expected_result"); er != "" {
		var expectedResult map[string]string
		json.Unmarshal([]byte(er), &expectedResult)
		job.ExpectedResult = expectedResult
	}
	bytes, _ := json.Marshal(job)
	resp, _, errs := gorequest.New().Post("http://localhost:3000/api/v1/jobs").Send(string(bytes)).End()

	if len(errs) != 0 {
		return cli.NewExitError(errs[0].Error(), 1)
	}
	if resp.StatusCode != 201 {
		return cli.NewExitError(fmt.Sprintf("Something went wrong, status: %d", resp.StatusCode), 1)
	}
	fmt.Println("Succesfully created a job")
	return nil
}
