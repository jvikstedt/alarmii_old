package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jvikstedt/alarmii/helper"
	"github.com/jvikstedt/alarmii/models"
	"github.com/jvikstedt/alarmii/scheduler"
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

	for running := true; running; {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Type q to quit")
		text, _ := reader.ReadString('\n')
		if text == "q\n" {
			running = false
		}
	}
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
