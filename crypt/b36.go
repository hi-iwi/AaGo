package crypt

const (
	Base36BinStr   = "0123456789abcdefghijklmnopqrstuvwxyz"
	ReverseBaseStr = byte('>') // 倒序
)

/*
Enigma：凯撒密码，简单替换密码；凯撒密码是把明文简单的平移 n 位，得到密文。
*/
func Enigma(v int, ps ...interface{}) byte {

	reverse := false
	asciiStart := int('a') - 10
	start := byte('0')

	for i := 0; i < len(ps); i++ {
		if r, ok := ps[i].(bool); ok {
			reverse = r
		} else if s, ok := ps[i].(byte); ok {
			start = s
		} else if s, ok := ps[i].(rune); ok {
			start = byte(s)
		}
	}

	if v > 35 {
		return byte('?')
	}
	// 87-96 97 - 122

	if start >= '0' && start <= '9' {
		start = 'a' - ('9' - start) - 1
	}
	var r int
	if !reverse {
		r = v + int(start)
		if r > int('z') {
			r = r - int('z') - 1 + asciiStart
		}
	} else {
		r = int(start) - v
		if r < asciiStart {
			r = int('z') - (asciiStart - r - 1)
		}
	}

	if r < int('a') {
		r = '9' - ('a' - r - 1)
	}
	return byte(r)
}

func B36(v int, ps ...interface{}) string {
	r := ""
	remainder := 0
	for {
		remainder = v % 36
		r = string(Enigma(remainder, ps...)) + r
		v /= 36
		if v <= 0 {
			break
		}
	}
	return r
}
