package aenum

// 配置文件协议类型
type ProtoType uint8

const (
	StringVtype ProtoType = 0
	// 为了兼容日后的 protobuf
	Int8Vtype    ProtoType = 1
	Uint8Vtype   ProtoType = 2
	Int16Vtype   ProtoType = 3
	Uint16Vtype  ProtoType = 4
	Int32Vtype   ProtoType = 5
	Uint32Vtype  ProtoType = 6
	Int64Vtype   ProtoType = 7
	Uint64Vtype  ProtoType = 8
	Float32Vtype ProtoType = 9
	Float64Vtype ProtoType = 10

	StringsVtype  ProtoType = 100 // 字符串数组
	Int8sVtype    ProtoType = 101
	Uint8sVtype   ProtoType = 102
	Int16sVtype   ProtoType = 103
	Uint16sVtype  ProtoType = 104
	Int32sVtype   ProtoType = 105
	Uint32sVtype  ProtoType = 106
	Int64sVtype   ProtoType = 107
	Uint64sVtype  ProtoType = 108
	Float32sVtype ProtoType = 109
	Float64sVtype ProtoType = 110
)
