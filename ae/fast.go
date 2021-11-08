package ae

// 服务端数据存储解析存在异常
func BadData(k, v string) *Error {
	return NewErr("Bad Data (%s:%s)", k, v)
}

//403 Forbidden
//服务器已经理解请求，但是拒绝执行它。与401响应不同的是，身份验证并不能提供任何帮助，而且这个请求也不应该被重复提交。如果这不是一个 HEAD 请求，而且服务器希望能够讲清楚为何请求不能被执行，那么就应该在实体内描述拒绝的原因。当然服务器也可以返回一个404响应，假如它不希望让客户端获得任何信息。
func Forbidden() *Error {
	return NewError(403, "Forbidden")
}

//409 Conflict
//由于和被请求的资源的当前状态之间存在冲突，请求无法完成。这个代码只允许用在这样的情况下才能被使用：用户被认为能够解决冲突，并且会重新提交新的请求。该响应应当包含足够的信息以便用户发现冲突的源头。
//冲突通常发生于对 PUT 请求的处理中。例如，在采用版本检查的环境下，某次 PUT 提交的对特定资源的修改请求所附带的版本信息与之前的某个（第三方）请求向冲突，那么此时服务器就应该返回一个409错误，告知用户请求无法完成。此时，响应实体中很可能会包含两个冲突版本之间的差异比较，以便用户重新提交归并以后的新版本。

func Conflict() *Error {
	return NewError(409, "Conflict")
}

func Expired() *Error {
	return NewError(410, "Expired")
}

func Gone() *Error {
	return NewError(410, "Gone")
}

func Locked() *Error {
	return NewError(423, "Locked")
}
