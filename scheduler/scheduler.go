package scheduler

import (
	"errors"

	"github.com/robfig/cron"
)

// Scheduleable interface for running schedules
type Scheduleable interface {
	Execute()
	CronTime() string
}

var c *cron.Cron

// SetupScheduler setups scheduler
func SetupScheduler() {
	c = cron.New()
}

// StartScheduler starts running scheduleables
func StartScheduler() {
	c.Start()
}

// StopScheduler stops running scheduleables
func StopScheduler() {
	c.Stop()
}

// AddScheduleable adds a scheduleable
func AddScheduleable(scheduleable Scheduleable) (err error) {
	cronTime := scheduleable.CronTime()
	if cronTime == "" {
		err = errors.New("CronTime not set")
	}
	err = c.AddFunc(cronTime, scheduleable.Execute)
	return
}

//func setupScheduler() {
//	c := cron.New()
//	projects, _ := models.GetProjects()
//	for _, p := range projects {
//		for _, j := range p.Jobs {
//			if j.Timing == "" {
//				continue
//			}
//			c.AddFunc(j.Timing, func() {
//				output, _ := exec.Command(j.Command, j.Arguments...).Output()
//				var objmap map[string]string
//				json.Unmarshal(output, &objmap)
//				for k, v := range j.ExpectedResult {
//					log.Info(objmap[k] + " == " + v)
//				}
//			})
//		}
//	}
//	c.Start()
//}
