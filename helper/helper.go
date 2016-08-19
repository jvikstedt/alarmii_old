package helper

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

// PIDFolder is path to pid folder
const PIDFolder = "tmp/pids"

// PIDName is pid file name
const PIDName = "alarmii.pid"

// Itob returns an 8-byte big endian representation of v.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// ReadPID returns main process pid
func ReadPID() (pid int, err error) {
	pidBytes, err := ioutil.ReadFile(PIDFolder + "/" + PIDName)
	if err != nil {
		return
	}
	pid, err = strconv.Atoi(string(pidBytes))
	return
}

// SavePID saves pid to tmp/pids folder
func SavePID() {
	pid := os.Getpid()
	err := os.MkdirAll(PIDFolder, 0755)
	if err != nil {
		panic(err)
	}
	s := strconv.Itoa(pid)
	err = ioutil.WriteFile(PIDFolder+"/"+PIDName, []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}

// EditFileWithDefaultEditor opens file with default editor
// If EDITOR env is not set, it uses vi
// Returns content of the file and error
func EditFileWithDefaultEditor(filePath string) (content []byte, err error) {
	defaultEditor := os.Getenv("EDITOR")
	if defaultEditor == "" {
		defaultEditor = "vi"
	}
	cmd := exec.Command(defaultEditor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return
	}
	content, err = ioutil.ReadFile(filePath)
	return
}
