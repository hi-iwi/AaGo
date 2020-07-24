package aa

import (
	"context"
	"fmt"
)

const TraceIdKey = "aa_trace_id"

func (app *Aa) Sprintf(ctx context.Context, msg string, args ...interface{}) string {
	msg = msg + " |<" + app.TraceID(ctx) + ">|"
	msg = fmt.Sprintf(msg, args...)
	return msg
}

func (app *Aa) TraceID(ctx context.Context) string {
	id, _ := ctx.Value(TraceIdKey).(string)
	return id
}

func (app *Aa) NewTraceableContext(trace string) context.Context {
	return context.WithValue(context.Background(), TraceIdKey, trace)
}
