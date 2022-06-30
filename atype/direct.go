package atype

func BoolUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
func IsTrue(u uint8) bool {
	return u > 0
}
func IsFalse(u uint8) bool {
	return u == 0
}
