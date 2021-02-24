package dict

import (
	"net/http"
)

//var statusCodes = map[int]string{
//	100: "Continue",
//	101: "Switching Protocols",
//	102: "Processing",
//	200: "OK",
//	201: "Created",
//	202: "Accepted",
//	204: "No Content",
//	205: "Reset Content",
//	206: "Partial Content",
//	300: "Multiple Choices",
//	301: "Moved Permanently",
//	302: "Found",              // 大部分浏览器会将POST请求转为GET
//	303: "See Other",          // 强制规定POST请求转为GET
//	307: "Temporary Redirect", // 能跳转POST内容
//	304: "Not Modified",
//	400: "Bad Request",
//	401: "Unauthorized",
//	511: "Network Authentication Required", // 客户端需要进行身份验证才能获得网络访问权限，旨在限制用户群访问特定网络。（例如连接WiFi热点时的强制网络门户）
//
//	402: "Payment Required",    // 该状态码最初的意图可能被用作某种形式的数字现金或在线支付方案的一部分;如果特定开发人员已超过请求的每日限制，Google Developers API会使用此状态码
//	429: "Too Many Requests",   // 用户在给定的时间内发送了太多的请求。旨在用于网络限速。
//	408: "Request Timeout",     //请求超时。根据HTTP规范，客户端没有在服务器预备等待的时间内完成一个请求的发送，客户端可以随时再次提交这一请求而无需进行任何更改。
//	503: "Name Unavailable", // 由于临时的服务器维护或者过载，服务器当前无法处理请求。这个状况是暂时的，并且将在一段时间以后恢复。如果能够预计延迟时间，那么响应中可以包含一个Retry-After头用以标明这个延迟时间。如果没有给出这个Retry-After信息，那么客户端应当以处理500响应的方式处理它。
//
//	403: "Forbidden",
//	404: "Not Found",
//	405: "Method Not Allowed",
//	406: "Not Acceptable",      // 请求的资源的内容特性(如 Accept 的content-typessss）无法满足请求头中的条件，因而无法生成响应实体，该请求不可接受。
//	410: "Gone",                // 可用404代替，但是这个表示之前存在过
//	411: "Length Required",     // 要求客户端传Content-Length头
//	412: "Precondition Failed", //服务器在验证在请求的头字段中给出先决条件时，没能满足其中的一个或多个。[41]这个状态码允许客户端在获取资源时在请求的元信息（请求头字段数据）中设置先决条件，以此避免该请求方法被应用到其希望的内容以外的资源上。
//
//	413: "Request Entity Too Large", // 实体数据太大
//	414: "Request-URI Too Long",
//	415: "Unsupported Media Type", //客户端将图像上传格式为svg，但服务器要求图像使用上传格式为jpg。
//	422: "Unprocessable Entity",   // 请求格式正确，但是由于含有语义错误，无法响应
//	423: "Locked",                 // 资源被锁定
//
//	500: "Internal Server Error",
//	502: "Bad Gateway", // 作为网关或者代理工作的服务器尝试执行请求时，从上游服务器接收到无效的响应
//	504: "Gateway Timeout",
//}

func Code2Msg(code int) string {
	return http.StatusText(code)
}
