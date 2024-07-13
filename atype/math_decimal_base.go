package atype

func formatWhole(s string, interval uint8) string {
	if interval == 0 || len(s) <= int(interval) {
		return s
	}
	var s2 string
	j := 0
	for i := len(s) - 1; i > -1; i-- {
		if j > 0 && j%int(interval) == 0 {
			s2 = "," + s2
		}
		s2 = string(s[i]) + s2
		j++
	}
	return s2
}
