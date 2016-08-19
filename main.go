package main

import (
	"github.com/jvikstedt/alarmii/cli"
	"github.com/jvikstedt/alarmii/helper"
)

func init() {
	helper.SetupLogger("log/info.log")
}

func main() {
	cli.Run()
}
