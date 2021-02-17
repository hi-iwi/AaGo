package code

import (
	"errors"
	"math"
	"strings"
)

var ObfuscateBase = []byte{'6', 'h', 'n', '1', '3', 'z', 's', 'm', 'c', 'd', 'o', 'f', 'i', 'j', '2', 'l', 'q', '0', '4', 'y', 'v', '5', '7', 'b', '9', 'a', 'x', 'g', 'w', 'k', 'u', '8', 't', 'p', 'e', 'r'}

func k2s(v int, shift int, base []byte) (s [2]byte) {
	l := len(base)
	shift %= l
	a := (v/l + shift) % l
	b := (v + shift) % l
	s[0] = base[a]
	s[1] = base[b]
	return
}
func s2k(s [2]byte, shift int, base []byte) int {
	l := len(base)
	shift %= l

	var (
		a, b int
	)
	for i, h := range base {
		if h == s[0] {
			a = i
		}
		if h == s[1] {
			b = i
		}
	}
	// 当发现解析后，跟之前结果不同。注意排查是不是 base 设置的不一样！
	x := ((a + l - shift) % l) * l
	y := (b + l - shift) % l
	return x + y
}

// @param shift，是 innerToken 第一个字符
// @return 结果一定是偶数，会自动填充的
func ObfuscateNumber(num uint64, shift int, base []byte) string {
	var b strings.Builder
	b.Grow(10)
	var x int
	var k [2]byte
	for {
		if num <= 0 {
			break
		}
		//用2个字符代表3个数字，倒序
		x = int(num % 1000)
		num = num / 1000
		k = k2s(x, shift, base)
		b.WriteByte(k[0])
		b.WriteByte(k[1])
	}

	return b.String()
}

func DeobfuscateNumber(u string, shift int, base []byte) (v uint64) {
	l := len(u)
	var s [2]byte
	var k float64

	if l%2 != 0 {
		return 0
	}
	// 一定是偶数
	for i := 0; i < l; i += 2 {
		s[0] = u[i]
		s[1] = u[i+1]
		k = float64(s2k(s, shift, base)) * math.Pow(1000, float64(i)/2.0)
		v += uint64(k)
	}
	return v
}

// 逐一字符混淆
func ObfuscateBytes(str []byte, shift int, base []byte) ([]byte, error) {
	var ok bool
	l := len(base)
	shift %= l
	for i, s := range str {
		ok = false
		for j, b := range base {
			if b == s {
				ok = true
				p := (j + shift) % l
				str[i] = base[p]
			}
		}
		if !ok {
			return nil, errors.New("all bytes must be contained in base")
		}
	}

	return str, nil
}

func DeobfuscateBytes(v []byte, shift int, base []byte) ([]byte, error) {
	var ok bool
	l := len(base)
	shift %= l
	for i, s := range v {
		ok = false
		for j, b := range base {
			if b == s {
				ok = true
				p := (j + l - shift) % l
				v[i] = base[p]
			}
		}
		if !ok {
			return nil, errors.New("all bytes must be contained in base")
		}
	}
	return v, nil
}
