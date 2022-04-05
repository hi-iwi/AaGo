package util

// 参考 strings.Index 写法
// 方便多字节字符查找 Index
func IndexRunes(s []rune, substr []rune) int {
	n := len(substr)
	switch {
	case n == 0:
		return 0
	case n == 1:
		for i, x := range s {
			if x == substr[0] {
				return i
			}
		}
		return -1
	}

	s0 := substr[0]
	m := len(s)
	for i, x := range s {
		if i+n > m {
			return -1
		}
		if x == s0 {
			var matched = true
			for j := 1; j < n; j++ {
				if s[i+j] != substr[j] {
					matched = false
					break
				}
			}
			if matched {
				return i
			}
		}
	}
	return -1
}

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
