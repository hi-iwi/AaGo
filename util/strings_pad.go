package util

// Pad 左填充字符串
// 也可能 pad 中文字符、特殊字符
// 返回字符串长度 >= minlen
func Pad(str, pad string, minlen int) string {
	if pad == "" || len(str) >= minlen {
		return str
	}
	for {
		str = pad + str
		if len(str) >= minlen {
			return str
		}
	}
}

// Unpad 移除左填充的字符串
func Unpad(str, pad string) string {
	n := len(str)
	m := len(pad)

	if n == 0 || m == 0 {
		return str
	}
	var x int
	for i := 0; i < n; i += m {
		if i+m > n-1 {
			break
		}
		if str[i:i+m] == pad {
			x = i + m
		}
	}
	if x > 0 {
		return str[x:]
	}
	return str
}
func PadRight(str string, pad string, minlen int) string {
	if pad == "" || len(str) >= minlen {
		return str
	}
	for {
		str += pad
		if len(str) >= minlen {
			return str
		}
	}
}
func UnpadRight(str, pad string) string {
	n := len(str)
	m := len(pad)

	if n == 0 || m == 0 {
		return str
	}
	var x int
	for i := n; i > -1; i -= m {
		if i-m < 0 {
			break
		}
		if str[i-m:i] == pad {
			x = i - m
		}
	}
	if x > 0 {
		return str[:x]
	}
	return str
}
