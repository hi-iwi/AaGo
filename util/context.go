package util

// CoContext a new context for coroutine
//func CoContext(ictx iris.Context) context.Context {
//	// @warn it must to be a child of context.Background(). Do not assign it to the child of iris.Context!
//	ctx := context.WithValue(context.Background(), alog.Sid, ictx.Values().GetString(string(alog.Sid))+"@co")
//	return ctx
//}
//
//func CoContextWithCancel(ictx iris.Context) (context.Context, context.CancelFunc) {
//	ctx, cancel := context.WithCancel(context.Background())
//	// @warn it must to be a child of context.Background(). Do not assign it to the child of iris.Context!
//	ctx = context.WithValue(ctx, alog.Sid, ictx.Values().GetString(string(alog.Sid))+"@co")
//	return ctx, cancel
//}
//
//func traceID(cs ...context.Context) string {
//	tid := ""
//	if len(cs) > 0 {
//		c := cs[0]
//		tid, _ = c.Value(alog.Sid).(string)
//	}
//	if tid == "" {
//		tid = uuid.New().Name()
//	}
//	return tid
//}
//func Ctx(cs ...context.Context) context.Context {
//	ctx := context.WithValue(context.Background(), alog.Sid, traceID(cs...)+"@co")
//	return ctx
//}
//
//func CtxWithCancel(cs ...context.Context) (context.Context, context.CancelFunc) {
//	ctx, cancel := context.WithCancel(context.Background())
//	ctx = context.WithValue(ctx, alog.Sid, traceID(cs...)+"@co")
//	return ctx, cancel
//}
