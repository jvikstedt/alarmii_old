package helper

import (
	"encoding/binary"
	"os"
	"strconv"
)

// Itob returns an 8-byte big endian representation of v.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// SavePID saves pid to tmp/pids folder
func SavePID() {
	pid := os.Getpid()
	err := os.MkdirAll("tmp/pids", 0755)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile("tmp/pids/alarmii.pid", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}
	s := strconv.Itoa(pid)
	f.Write([]byte(s))
	f.Close()
}
