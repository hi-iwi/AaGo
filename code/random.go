package code

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Random 返回随机生成的字符串
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
// @url https://www.crockford.com/base32.html
func Random(strlen int, base int, humanReadable bool) string {
	var (
		seed int64
		num  int
		s    string
	)
	b := strings.Builder{}
	b.Grow(strlen)

	for i := 0; i < strlen; i++ {
		seed = time.Now().UnixNano() + int64(i+num)
		num = rand.New(rand.NewSource(seed)).Intn(base)
		if base == 10 {
			s = strconv.Itoa(num)
		} else {
			s = strconv.FormatInt(int64(num), base)
			if base > 16 && humanReadable {
				switch s {
				case "o":
					s = "0"
				case "i", "l":
					s = "1"
				case "u":
					s = "v"
				case "z":
					s = "2"
				}
			}
		}
		b.WriteString(s)
	}

	return b.String()
}
