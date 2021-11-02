package code

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

// crypto/rand是为了提供更好的随机性满足密码对随机数的要求，在linux上已经有一个实现就是/dev/urandom，crypto/rand 就是从这个地方读“真随机”数字返回，但性能比较慢
// 每次更新文件的时候，随机数就会变化；如果不更新文件，就是伪随机数（不停重复生成）
func RandNum(len int) (string, error) {
	x := int64(math.Pow10(len)) - 1
	n, err := rand.Int(rand.Reader, big.NewInt(x))
	if err != nil {
		return "", err
	}

	str := fmt.Sprintf("%0"+strconv.Itoa(len)+"s", n.String())
	return str, nil
}

// Random 返回随机生成的字符串
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
// @url https://www.crockford.com/base32.html
//func Random(strlen int, base int, humanReadable bool) string {
//	var (
//		seed int64
//		num  int
//		s    string
//	)
//	b := strings.Builder{}
//	b.Grow(strlen)
//
//	for i := 0; i < strlen; i++ {
//		seed = time.Now().UnixNano() + int64(i+num)
//		num = rand.New(rand.NewSource(seed)).Intn(base)
//		if base == 10 {
//			s = strconv.Itoa(num)
//		} else {
//			s = strconv.FormatInt(int64(num), base)
//			if base > 16 && humanReadable {
//				switch s {
//				case "o":
//					s = "0"
//				case "i", "l":
//					s = "1"
//				case "u":
//					s = "v"
//				case "z":
//					s = "2"
//				}
//			}
//		}
//		b.WriteString(s)
//	}
//
//	return b.String()
//}
