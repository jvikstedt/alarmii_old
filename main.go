package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

func setupLogger() {
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  "log/info.log",
		log.ErrorLevel: "log/error.log",
	}))
}

func init() {
	setupLogger()
}

func main() {
	app := cli.NewApp()
	app.Name = "Alarmii"
	app.Usage = ""
	app.Action = func(c *cli.Context) error {
		fmt.Println("Test")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "Prints config",
			Action: func(c *cli.Context) error {
				config := LoadConfig("config.json")
				log.Info(config)
				return nil
			},
		},
	}

	app.Run(os.Args)

	//config := LoadConfig("config.json")
	//output, _ := exec.Command(config.Projects[0].Jobs[0].Command, config.Projects[0].Jobs[0].Arguments...).Output()
}
