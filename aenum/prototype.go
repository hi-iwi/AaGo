package aenum

import "strconv"

// 配置文件协议类型
type ProtoType uint8

const (
	Nil     ProtoType = 0
	Bool    ProtoType = 1
	Booln   ProtoType = 2 // 1|0
	String  ProtoType = 3
	Float64 ProtoType = 4
	Float32 ProtoType = 5

	Uint64 ProtoType = 10
	Uint32 ProtoType = 11
	Uint24 ProtoType = 12
	Uint16 ProtoType = 13
	Uint8  ProtoType = 14

	Int64 ProtoType = 15
	Int32 ProtoType = 16
	Int24 ProtoType = 17
	Int16 ProtoType = 18
	Int8  ProtoType = 19

	Bools    ProtoType = 101 // bool array
	Boolns   ProtoType = 102
	Strings  ProtoType = 103
	Float64s ProtoType = 104
	Float32s ProtoType = 105

	Uint64s ProtoType = 110
	Uint32s ProtoType = 111
	Uint24s ProtoType = 112
	Uint16s ProtoType = 113
	Uint8s  ProtoType = 114
	Int64s  ProtoType = 115
	Int32s  ProtoType = 116
	Int24s  ProtoType = 117
	Int16s  ProtoType = 118
	Int8s   ProtoType = 119

	Date      ProtoType = 120
	Time      ProtoType = 121
	Datetime  ProtoType = 122
	UnixTime  ProtoType = 123
	Year      ProtoType = 124
	YearMonth ProtoType = 125 // uint24 date: yyyymm  不要用 Date，主要是不需要显示dd。

	Money ProtoType = 130
	Price ProtoType = 131

	Distri   ProtoType = 140 // 6 位地址简码
	AddrId   ProtoType = 141 // 12 位地址码
	CountryT ProtoType = 142

	Struct ProtoType = 255
)

func NewProtoType(t uint8) (ProtoType, bool) {
	return ProtoType(t), true
}
func (x ProtoType) Uint8() uint8         { return uint8(x) }
func (x ProtoType) String() string       { return strconv.FormatUint(uint64(x), 10) }
func (x ProtoType) Is(x2 ProtoType) bool { return x == x2 }
func (x ProtoType) In(args ...ProtoType) bool {
	for _, a := range args {
		if a == x {
			return true
		}
	}
	return false
}
