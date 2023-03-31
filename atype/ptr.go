package atype

func PtrUint64(n *uint64) uint64 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint32(n *uint32) uint32 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint(n *uint) uint {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint24(n *Uint24) Uint24 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint16(n *uint16) uint16 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrUint8(n *uint8) uint8 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt64(n *int64) int64 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt32(n *int32) int32 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt(n *int) int {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt16(n *int16) int16 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrInt8(n *int8) int8 {
	if n == nil {
		return 0
	}
	return *n
}
func PtrFloat64(n *float64) float64 {
	if n == nil {
		return 0.00
	}
	return *n
}
func PtrFloat32(n *float32) float32 {
	if n == nil {
		return 0.00
	}
	return *n
}
