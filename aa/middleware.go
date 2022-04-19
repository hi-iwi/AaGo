package aa

import (
	"context"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

const ErrorLevelKey = "aa_error_level"
const TraceIdKey = "aa_trace_id"

// tracePrefix should be captialized
func (app *Aa) IrisMiddleware(ictx iris.Context) {
	defer ictx.Next() // 这个是必须要存在的！！！
	traceId := ictx.GetHeader("X-Request-Wid")
	if traceId == "" {
		traceId = uuid.New().String()
	}
	ictx.Values().Set(TraceIdKey, traceId)
}

// 这里会整体clone一个context，性能并不好。但是为了代码美化，还是牺牲这点性能，换取统一的trace id传递
func Context(ictx iris.Context) context.Context {
	traceId := ictx.Values().GetString(TraceIdKey)
	return context.WithValue(ictx.Request().Context(), TraceIdKey, traceId)
}

func ContextWithTraceID(ctx context.Context, traceId string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, TraceIdKey, traceId)
}

// 使用 context.WithValue 会复制整个 context，会比较慢。尽量直接用 ictx.Values()
//func SprintfTrace(ctx context.Context, msg string, args ...interface{}) string {
//	msg = fmt.Sprintf(msg, args...)
//	return msg + " " + Sid(ctx)
//}

func TraceID(ctx context.Context) string {
	traceId, ok := ctx.Value(TraceIdKey).(string)
	if ok {
		return traceId
	}
	return "no trace id"
}
