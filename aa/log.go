package aa

import (
	"sync"
)

const (
	EnvProd = uint8(1)
	EnvDev  = uint8(9)

	LogDebug = iota
	LogInfo
	LogNotice
	LogWarning
	LogErr
	LogCrit
	LogAlert
	LogEmerge
)

var (
	env        = EnvProd
	logger     *Log
	newLogOnce sync.Once
)

type Log struct {
	priority int
}

func NewLog(p int) *Log {
	newLogOnce.Do(func() {
		logger = &Log{
			priority: p,
		}
	})
	return logger
}
