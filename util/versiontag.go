package util

import (
	"os"
	"os/exec"
)

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
