package aa

import (
	"context"
	"github.com/hi-iwi/AaGo/ae"
	"runtime"
	"strconv"
	"strings"
)



type Log interface {
	AError(ctx context.Context, e *ae.Error)
	// AuthDebug 包含详细的开发情报的信息，通常只在调试一个程序时使用
	Debug(ctx context.Context, msg string, args ...interface{})

	// Info 情报信息，正常的系统消息，比如骚扰报告，带宽数据等，不需要处理。
	Info(ctx context.Context, msg string, args ...interface{})

	// Notice 不是错误情况，也不需要立即处理。
	Notice(ctx context.Context, msg string, args ...interface{})

	// Warn 警告信息，不是错误，比如系统磁盘使用了85%等。
	Warn(ctx context.Context, msg string, args ...interface{})

	// Error 错误，不是非常紧急，在一定时间内修复即可。
	Error(ctx context.Context, msg string, args ...interface{})

	// Crit 重要情况，如硬盘错误，备用连接丢失
	Crit(ctx context.Context, msg string, args ...interface{})

	// Alert 应该被立即改正的问题，如系统数据库被破坏，ISP连接丢失。
	Alert(ctx context.Context, msg string, args ...interface{})

	// Emerg 紧急情况，需要立即通知技术人员。
	Emerg(ctx context.Context, msg string, args ...interface{})

	Printf(ctx context.Context, msg string, args ...interface{})

	Println(ctx context.Context, msg ...interface{})

	// Trace 跟踪请求链路，用于性能监控
	Trace(ctx context.Context)
}

func caller(skip int) string {
	pc, file, line, _ := runtime.Caller(skip)
	pcs := runtime.FuncForPC(pc).Name() // 函数名
	a := strings.Split(file, "/")       // 文件名
	return a[len(a)-1] + ":" + strconv.Itoa(line) + " " + pcs
}

func traceid(ctx context.Context) string {
	id, _ := ctx.Value(TraceIdKey).(string)
	return id
}
