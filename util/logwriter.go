package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

var onceLogWriter sync.Once
var logWriter *LogWriter

type LogWriter struct {
	dir     string
	file    *os.File
	logName string
	sync    sync.Mutex
}

func (lw *LogWriter) Write(p []byte) (n int, err error) {

	tm := time.Now()
	newLogFile, _ := filepath.Abs(path.Join(path.Dir(lw.app.Config.Get("app.log_path").String()), tm.Format("2006-01-02")+".bak.log"))
	file := lw.file
	var linkName string
	lw.sync.Lock()
	if newLogFile != lw.logName {
		f, err := os.OpenFile(newLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err == nil {
			f.WriteString("\n\n\n")
			lw.logName = newLogFile
			if lw.file != nil {
				lw.file.Close()
			}
			lw.file = f
			file = f
			linkName, _ = filepath.Abs(lw.app.Config.Get("app.log_path").String())
		} else {
			fmt.Println(err)
		}
	}
	lw.sync.Unlock()

	var tim time.Time
	var oldLogName string
	for i := 90; i < 180; i++ {
		tim = tm.Add(-time.Hour * 24 * time.Duration(i))
		oldLogName, _ = filepath.Abs(path.Join(path.Dir(lw.app.Config.Get("app.log_path").String()), tim.Format("2006-01-02")+".bak.log"))
		os.Remove(oldLogName)
	}

	if linkName != "" {
		os.Remove(linkName)
		os.Symlink(newLogFile, linkName)
	}
	if file != nil {
		return file.Write(p)
	}
	return 0, nil
}

func NewLogWriter(dir string) *LogWriter {

	onceLogWriter.Do(func() {
		logWriter = &LogWriter{
			dir: dir,
		}
	})
	return logWriter
}
