package aenum

// func NewDemo() (Demo, bool){}
type Demo interface {
	Valid() bool                  // 是否在正确范围内
	//Raw()  interface{}           // 原始值，如 uint16、byte、string等
	String() string               // 类型转换为string
	Name() string                 // 显示给客户端，同时为了避免客户端侦测太多服务端信息，特别是作为加密因子
	Is(demo Demo) bool // 判断是否相等
	In(series []interface{}) bool // 是否在某个系列里面
}
