package ae

/*
自定义的：444 NoRows； 503
*/
var (
	// @warn 自定义。444 空数组返回这个错误，表示不可以再进行下一页查询了
	// 若是 code:204, data:[]  空数组，表示查询到了数据，但是数据过滤完了，可以尝试下一页查询
	NoContent       = &Error{204, "No Content"}
	Unauthorized    = &Error{401, "Unauthorized"}
	PaymentRequired = &Error{402, "Payment Required"}
	//403 Forbidden
	//服务器已经理解请求，但是拒绝执行它。与401响应不同的是，身份验证并不能提供任何帮助，而且这个请求也不应该被重复提交。如果这不是一个 HEAD 请求，而且服务器希望能够讲清楚为何请求不能被执行，那么就应该在实体内描述拒绝的原因。当然服务器也可以返回一个404响应，假如它不希望让客户端获得任何信息。
	Forbidden = &Error{403, "Forbidden"}
	NotFound  = &Error{404, "Not Found"} // refer to redis.Nil, sql.ErrNoRows
	// 空数组返回 422
	//409 Conflict
	//由于和被请求的资源的当前状态之间存在冲突，请求无法完成。这个代码只允许用在这样的情况下才能被使用：用户被认为能够解决冲突，并且会重新提交新的请求。该响应应当包含足够的信息以便用户发现冲突的源头。
	//冲突通常发生于对 PUT 请求的处理中。例如，在采用版本检查的环境下，某次 PUT 提交的对特定资源的修改请求所附带的版本信息与之前的某个（第三方）请求向冲突，那么此时服务器就应该返回一个409错误，告知用户请求无法完成。此时，响应实体中很可能会包含两个冲突版本之间的差异比较，以便用户重新提交归并以后的新版本。
	Timeout = &Error{408, "Request Timeout"} // 被限流也是返回这个

	Conflict           = &Error{409, "Conflict"}
	Gone               = &Error{410, "Gone"} // 以前存在过，以后都不会再存在了，表示数据已经删除、过期、失效
	PreconditionFailed = &Error{412, "Precondition Failed"}
	BadMediaType       = &Error{415, "Unsupported Media Type"} // 上传的数据格式非法

	// @warn 自定义。444 空数组返回这个错误，表示不可以再进行下一页查询了
	// 若是 code:204, data:[]  空数组，表示查询到了数据，但是数据过滤完了，可以尝试下一页查询
	NoRows = &Error{444, "No Rows"}

	//RetryWith = &Error{449, "Locked"}

	Locked           = &Error{423, "Locked"}
	FailedDependency = &Error{424, "Failed Dependency"} // 之前发生错误
	TooEarly         = &Error{425, "Too Early"}         // 表示服务器不愿意冒险处理可能被重播的请求。
	TooManyRequests  = &Error{429, "Too Many Requests"} // 用户在给定的时间内发送了太多请求（"限制请求速率"）

	Illegal = &Error{451, "Unavailable For Legal Reasons"} // 该请求因政策法律原因不可用。

	InternalServerError    = &Error{500, "Internal Server Error"}
	NotImplemented         = &Error{501, "Not Implemented"}          // 服务器不支持当前请求所需要的某个功能。当服务器无法识别请求的方法，并且无法支持其对任何资源的请求。
	BadGateway             = &Error{502, "Bad Gateway"}              // 上游服务错误
	ServerException        = &Error{503, "Server Exception"}         // 客户端自定义，表示未知服务端错误；最常见的就是，没有正确返回数据，或者返回 {code:0,msg:""} 等未协商的数据，导致客户端无法正常处理
	GatewayTimeout         = &Error{504, "Gateway Timeout"}          // 上游服务器超时
	BandwidthLimitExceeded = &Error{509, "Bandwidth Limit Exceeded"} // 限流
	StatusException        = &Error{555, "Server Status Exception"}  // http 状态码出错，未达到程序阶段
)

func RetryWith(redirect string) *Error {
	return &Error{449, redirect} // 特殊错误码，msg 用于跳转
}

// 服务端数据存储解析存在异常
func ParseFailed(k, v string) *Error {
	return NewErr("parse failed (%s:%s)", k, v)
}
func BadParam(param string) *Error {
	return NewError(400, "bad parameter `"+param+"`")
}
func ProxyAuthRequired(msg string) *Error {
	return NewError(407, "proxy auth required: "+msg)
}
func Catch(es ...*Error) *Error {
	for _, e := range es {
		if e != nil {
			return e
		}
	}
	return nil
}
func CatchError(es ...error) error {
	for _, e := range es {
		if e != nil {
			return e
		}
	}
	return nil
}
