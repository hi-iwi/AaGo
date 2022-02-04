package dtype

func Bool2Uint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
func Uint82Bool(u uint8) bool {
	return u > 0
}
