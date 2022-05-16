package aa

import (
	"context"
	"github.com/hi-iwi/AaGo/ae"
)

type ErrorLevel uint8

const (
	AllError ErrorLevel = iota
	Debug
	Info
	Notice
	Warn
	Err
	Crit
	Alert
	Emerg
)

func (lvl ErrorLevel) Name() string {
	switch lvl {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Notice:
		return "notice"
	case Warn:
		return "warn"
	case Err:
		return "err"
	case Crit:
		return "crit"
	case Alert:
		return "alert"
	case Emerg:
		return "emerg"
	}
	return ""
}

type Log interface {
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

	Println(ctx context.Context, msg ...interface{})

	// Trace 跟踪请求链路，用于性能监控
	Trace(ctx context.Context)
}

func traceid(ctx context.Context) string {
	id, _ := ctx.Value(TraceIdKey).(string)
	return id
}
func errorlevel(ctx context.Context) ErrorLevel {
	level, _ := ctx.Value(ErrorLevelKey).(ErrorLevel)
	return level
}

// 快捷方式，对服务器错误记录日志
func (app *App) Try(ctx context.Context, e *ae.Error) bool {
	if e != nil && e.IsServerError() {
		app.Log.Error(ctx, e.Error())
		return false
	}
	return true
}

// 快捷记录错误
func (app *App) TryLog(ctx context.Context, err error) {
	if err != nil {
		app.Log.Error(ctx, ae.Caller(1)+" "+err.Error())
	}
}

// 快捷panic
func (app *App) TryPanic(ctx context.Context, e *ae.Error) {
	if e != nil {
		app.Log.Error(ctx, ae.Caller(1)+" "+e.Error())
		panic(e.Error())
	}
}
