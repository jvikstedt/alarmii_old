package scheduler

import (
	"errors"

	"github.com/robfig/cron"
)

// Schedulable interface that can be used to start a new schedulable
type Schedulable interface {
	Run()
	CronTime() string
}

var c *cron.Cron

// SetupScheduler setups scheduler
func SetupScheduler() {
	c = cron.New()
}

// StartScheduler starts running jobs
func StartScheduler() {
	c.Start()
}

// StopScheduler stops running jobs
func StopScheduler() {
	c.Stop()
}

// AddSchedulable adds a schedulable
func AddSchedulable(schedulable Schedulable) (err error) {
	if time := schedulable.CronTime(); time == "" {
		err = errors.New("schedulable.CronTime() returned an empty string")
	} else {
		err = c.AddJob(time, schedulable)
	}
	return
}
