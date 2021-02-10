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
func xlogHeader(ctx context.Context, caller string, severity string) string {
	tid := traceid(ctx)
	b := strings.Builder{}
	b.Grow(15 + len(tid))
	b.WriteString(caller)
	if tid != "" {
		b.WriteString("{" + tid + "}")
	}
	if severity != "" {
		b.WriteString("[" + severity + "]")
	}
	return b.String()
}

func xprintf(ctx context.Context, skip int, severity string, msg string, args ...interface{}) {
	head := xlogHeader(ctx, caller(skip), severity)
	msg = head + msg
	if len(args) == 0 {
		log.Println(msg)
	} else {
		log.Printf(msg+"\n", args...)
	}
}

func (l *xlog) Debug(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "debug", msg, args...)
}

func (l *xlog) Info(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "info", msg, args...)
}

func (l *xlog) Notice(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "notice", msg, args...)
}

func (l *xlog) Warn(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "warn", msg, args...)
}

func (l *xlog) Error(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "error", msg, args...)
}

func (l *xlog) Crit(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "crit", msg, args...)
}

func (l *xlog) Alert(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "alert", msg, args...)
}

func (l *xlog) Emerg(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "error", msg, args...)
}

func (l *xlog) Printf(ctx context.Context, msg string, args ...interface{}) {
	xprintf(ctx, 3, "", msg, args...)
}

func (l *xlog) Println(ctx context.Context, msg ...interface{}) {
	log.Println(xlogHeader(ctx, caller(2), ""), fmt.Sprint(msg...))
}

func (l *xlog) Trace(ctx context.Context) {
	tid := traceid(ctx)
	log.Printf("{%s} %s\n", tid, caller(2))
}

func (l *xlog) AError(ctx context.Context, e *ae.Error) {
	if e.Code >= 500 {
		xprintf(ctx, 3, "", e.Error())
	}
}
