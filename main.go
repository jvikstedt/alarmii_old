package main

import "github.com/jvikstedt/alarmii/logger"

func init() {
	logger.Setup("development.log")
}

func main() {
	logger.Cleanup()
}
