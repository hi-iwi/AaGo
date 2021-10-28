package ai

import "strconv"

type (
	// 这里是放在 DTO 里面的，而不是放在 entity 里面的
	Uint64Str string // Id 数据存的是 uint64，拿取到接口都是 string
	Uint32Str string // uint/uint32
	Uint16Str string
	Uint8Str  string
)

func ToUint64Str(id uint64) Uint64Str {
	return Uint64Str(strconv.FormatUint(id, 10))
}

func ToUint32Str(id uint32) Uint32Str {
	return Uint32Str(strconv.FormatUint(uint64(id), 10))
}
func ToUint16Str(id uint16) Uint16Str {
	return Uint16Str(strconv.FormatUint(uint64(id), 10))
}

func ToUint8Str(id uint8) Uint8Str {
	return Uint8Str(strconv.FormatUint(uint64(id), 10))
}
