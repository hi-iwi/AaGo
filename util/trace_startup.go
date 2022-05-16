package util

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync/atomic"
	"time"
)

var startingSteps int32
var gitHash string

func GitVersion() string {
	if gitHash != "" {
		return gitHash
	}
	s, _ := exec.LookPath(os.Args[0])
	file, err := os.OpenFile(s, os.O_RDONLY, 0666)
	if err != nil {
		return ""
	}
	defer file.Close()
	finfo, _ := file.Stat()
	hash := make([]byte, 40)
	file.ReadAt(hash, finfo.Size()-40-1)
	// git log id (or hash) only contains number and lower case alphabet
	for _, h := range hash {
		if h < '0' || (h > '9' && h < 'a') || h > 'z' {
			return ""
		}
	}
	gitHash = string(hash)
	return gitHash
}
func TraceStartup(msg ...string) {
	id := atomic.AddInt32(&startingSteps, 1)
	//log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate | log.Lmicroseconds)
	n := time.Now()
	now := n.Format("2006-01-02 15:04:05") + "." + strconv.FormatInt(n.UnixMicro(), 10)
	m := now + " starting " + strconv.FormatInt(int64(id), 10)
	if len(msg) > 0 {
		m += " " + msg[0]
	}
	fmt.Println(m)
}
