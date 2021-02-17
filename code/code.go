package code

import "strings"

func TrimPad(s string, padStr string, fixlen int) string {
	if len(s) >= fixlen {
		return s[0:fixlen]
	}
	return Pad(s, padStr, fixlen)
}
func Pad(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
