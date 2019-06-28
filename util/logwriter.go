package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

var onceLogWriter sync.Once
var logWriter *LogWriter

type LogWriter struct {
	filename string
	file     *os.File
	logName  string
	sync     sync.Mutex
}

func RedirectLog(logfile string, crashfile string, mode os.FileMode) {
	os.MkdirAll(path.Dir(logfile), mode)
	os.MkdirAll(path.Dir(crashfile), mode)

	logfile, _ = filepath.Abs(logfile)

	file, err := os.OpenFile(crashfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err == nil {
		syscall.Dup2(int(file.Fd()), 2)
	}

	log.SetOutput(NewLogWriter(logfile))
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate | log.Lmicroseconds)
}

func (lw *LogWriter) Write(p []byte) (n int, err error) {
	dir := path.Dir(lw.filename)
	tm := time.Now()
	newLogFile := path.Join(dir, tm.Format("2006-01-02")+".bak.log")
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
			linkName = lw.filename
		} else {
			fmt.Println(err)
		}
	}
	lw.sync.Unlock()

	var tim time.Time
	var oldLogName string
	for i := 90; i < 180; i++ {
		tim = tm.Add(-time.Hour * 24 * time.Duration(i))
		oldLogName = path.Join(dir, tim.Format("2006-01-02")+".bak.log")
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

func NewLogWriter(logfile string) *LogWriter {

	onceLogWriter.Do(func() {
		logWriter = &LogWriter{
			filename: logfile,
		}
	})
	return logWriter
}
