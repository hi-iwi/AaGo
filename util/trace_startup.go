package util

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync/atomic"
)

var startingSteps int32
var gitHash string

func GitVersion() string {
	if gitHash == "" {
		s, _ := exec.LookPath(os.Args[0])
		file, err := os.OpenFile(s, os.O_RDONLY, 0666)
		if err == nil {
			defer file.Close()
			finfo, _ := file.Stat()
			hash := make([]byte, 40)
			file.ReadAt(hash, finfo.Size()-40-1)
			gitHash = string(hash)
		}
	}
	return gitHash
}
func TraceStartup(msg ...string) {
	id := atomic.AddInt32(&startingSteps, 1)
	m := "starting " + strconv.FormatInt(int64(id), 10)
	if len(msg) > 0 {
		m += " " + msg[0]
	}
	log.Println(msg)
}
