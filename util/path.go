package util

import "strings"

// similar to path.Base()
func Filename(p string) string {
	if p == "" {
		return ""
	}
	i := strings.LastIndexByte(p, '/')
	if i == len(p) {
		return ""
	}
	return p[i+1:]
}
