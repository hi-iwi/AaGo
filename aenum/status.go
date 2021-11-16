package aenum

type Status int8

type DataParsing int8 // 解析远程数据，储存远程数据记录时用到

const (
	Deleted       Status = -128 //已删除
	PendingReview Status = -1   // 未审核，不展示
	Created       Status = 0    // 显示
	Verified      Status = 1    // 审核通过，展示
	Topping       Status = 100  // 未完成，置顶展示
	Closed        Status = 126  // 已关闭，展示
	Finished      Status = 127  // 已完成，展示
)



const (
	DataParsingCheckFailed DataParsing = -2 // 数据签名核对错误、字段核对错误
	DataParsingFailed      DataParsing = -1 // 数据解析失败
	DataParsingBizFailed   DataParsing = 0  // 数据解析成功了，但是业务结果返回失败
	DataParsingBizOK       DataParsing = 1  // 数据解析成功了，并且业务结果返回成功
)

const (
	HttpStatusAccountConflict = 490 // 授权登录成功，但是该授权账户绑定过的UID，和当前登录的UID不一致。
	HttpStatusAccountUnlinked = 491 // 已经授权登陆，但是需要绑定手机号/账号
)


