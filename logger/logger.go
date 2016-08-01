package logger

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
)

var file *os.File

// Setup sets up log file
func Setup(fileName string) {
	if err := os.MkdirAll("log", 0777); err != nil {
		panic(err)
	}
	file, err := os.OpenFile(fmt.Sprintf("log/%s", fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
}

// Cleanup closes log file
func Cleanup() {
	file.Close()
}

// Info info level log
func Info(args ...interface{}) {
	logrus.Info(args)
}

// Warn warn level log
func Warn(args ...interface{}) {
	logrus.Warn(args)
}

// Error error level log
func Error(args ...interface{}) {
	logrus.Error(args)
}
