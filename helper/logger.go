package helper

import (
	"log"
	"os"
)

type Logger struct {
	FilePath string
}

func SetupLogger(filePath string) {
	logger := Logger{FilePath: filePath}
	log.SetOutput(logger)
}

func (l Logger) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile("log/info.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	n, err = f.Write(p)
	return
}
