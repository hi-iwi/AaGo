package aenum

type DataParsing int8 // 解析远程数据，储存远程数据记录时用到
const (
	DataParsingCheckFailed DataParsing = -2 // 数据签名核对错误、字段核对错误
	DataParsingFailed      DataParsing = -1 // 数据解析失败
	DataParsingBizFailed   DataParsing = 0  // 数据解析成功了，但是业务结果返回失败
	DataParsingBizOK       DataParsing = 1  // 数据解析成功了，并且业务结果返回成功
)
