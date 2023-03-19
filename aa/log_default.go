package aa

import (
	"context"
	"fmt"
	"github.com/hi-iwi/AaGo/ae"
	"log"
	"strings"
	"sync"
)

type xlog struct {
}

var (
	xlogOnce     sync.Once
	xlogInstance *xlog
)

func NewDefaultLog() Log {
	xlogOnce.Do(func() {
		xlogInstance = &xlog{}
	})
	return xlogInstance
}
func xlogHeader(ctx context.Context, caller string, level ErrorLevel) string {
	tid := traceid(ctx)
	b := strings.Builder{}
	b.Grow(15 + len(tid))
	b.WriteString(caller)
	if tid != "" {
		b.WriteString("{" + tid + "}")
	}
	if level != AllError {
		b.WriteString("[" + level.Name() + "]")
	}
	return b.String()
}

func xprintf(ctx context.Context, skip int, level ErrorLevel, msg string, args ...interface{}) {
	errlevel := errorlevel(ctx)
	if errlevel != AllError && errlevel&level == 0 {
		return
	}
	skip++ // 跳出 xprintf 自身函数
	head := xlogHeader(ctx, ae.Caller(skip), level)
	msg = head + msg
	if len(args) == 0 {
		log.Println(msg)
	} else {
		log.Printf(msg+"\n", args...)
	}
}
func (l *xlog) New(prefix string, f func(context.Context, int, string, ...interface{}), suffix ...string) func(context.Context, int, string, ...interface{}) {
	var s string
	if len(suffix) > 0 {
		s = " " + suffix[0]
	}
	if prefix != "" {
		prefix += " "
	}
	return func(ctx context.Context, skip int, msg string, args ...interface{}) {
		f(ctx, skip+2, prefix+msg+s, args...)
	}
}
func (l *xlog) Debug(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Debug, msg, args...)
}

func (l *xlog) Info(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Info, msg, args...)
}

func (l *xlog) Notice(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Notice, msg, args...)
}

func (l *xlog) Warn(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Warn, msg, args...)
}

func (l *xlog) Error(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Err, msg, args...)
}

func (l *xlog) Crit(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Crit, msg, args...)
}

func (l *xlog) Alert(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Alert, msg, args...)
}

func (l *xlog) Emerg(ctx context.Context, skip int, msg string, args ...interface{}) {
	xprintf(ctx, skip+1, Emerg, msg, args...)
}

func (l *xlog) Println(ctx context.Context, skip int, msg ...interface{}) {
	log.Println(xlogHeader(ctx, ae.Caller(skip+1), AllError), fmt.Sprint(msg...))
}

func (l *xlog) Trace(ctx context.Context, skip int) {
	tid := traceid(ctx)
	log.Printf("[TRACE]{%s} %s\n", tid, ae.Caller(skip+1))
}
