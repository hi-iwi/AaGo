package aa

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
)

const ErrorLevelKey = "aa_error_level"
const CtxParamTraceKey = "Trace"
const IctxParamTraceId = "TraceId" // nginx 层传递过来
const IctxParamRemoteAddress = "RemoteAddr"
const IctxParamVuser = "TraceVuser"

// tracePrefix should be captialized
func (app *App) IrisMiddleware(ictx iris.Context) {
	defer ictx.Next() // 这个是必须要存在的！！！
	ictx.Values().Set(IctxParamRemoteAddress, ictx.RemoteAddr())
}

func SetIctxTraceInfoVuser(ictx iris.Context, vuser string) {
	ictx.Values().Set(IctxParamVuser, vuser)
}

// 这里会整体clone一个context，性能并不好。但是为了代码美化，还是牺牲这点性能，换取统一的trace id传递
func Context(ictx iris.Context) context.Context {
	traceId := TraceId(ictx)
	ip := RemoteAddr(ictx)
	vuser := ictx.Values().GetString(IctxParamVuser)
	var trace string
	if vuser == "" {
		trace = fmt.Sprintf("{%s %s}", traceId, ip)
	} else {
		trace = fmt.Sprintf("{%s %s %s}", traceId, ip, vuser)
	}
	return context.WithValue(ictx.Request().Context(), CtxParamTraceKey, trace)
}
func traceInfo(ctx context.Context) string {
	trace, _ := ctx.Value(CtxParamTraceKey).(string)
	return trace
}

// 一般用于任务，因此没有客户端，无法获取到IP、TraceId等
func ContextWithTraceID(ctx context.Context, traceId string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, CtxParamTraceKey, "{"+traceId+"}")
}

// 使用 context.WithValue 会复制整个 context，会比较慢。尽量直接用 ictx.Values()
//func SprintfTrace(ctx context.Context, msg string, args ...any) string {
//	msg = fmt.Sprintf(msg, args...)
//	return msg + " " + Sid(ctx)
//}

func TraceId(ictx iris.Context) string {
	traceId := ictx.Values().GetString(IctxParamTraceId)
	if traceId != "" {
		return traceId
	}
	traceId = ictx.GetHeader("X-TraceId")
	if traceId != "" {
		ictx.Values().Set(IctxParamTraceId, traceId)
	}
	return traceId
}
func RemoteAddr(ictx iris.Context) string {
	ip := ictx.Values().GetString(IctxParamRemoteAddress)
	if ip != "" {
		return ip
	}
	ip = ictx.RemoteAddr()
	if ip != "" {
		ictx.Values().Set(IctxParamRemoteAddress, ip)
	}
	return ip
}
