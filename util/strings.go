package util

// 也可能 pad 中文字符、特殊字符
func Pad(str string, pad string, minlen int) string {
	if len(str) >= minlen {
		return str
	}
	for {
		str = pad + str
		if len(str) >= minlen {
			return str
		}
	}
}

func PadRight(str string, pad string, minlen int) string {
	if len(str) >= minlen {
		return str
	}
	for {
		str += pad
		if len(str) >= minlen {
			return str
		}
	}
}
