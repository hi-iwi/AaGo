package crypt

import "strconv"

func Pad(n interface{}, l int) string {
	s := ""
	if a, ok := n.(int); ok {
		s = strconv.Itoa(a)
	} else if b, ok := n.(string); ok {
		s = b
	}
	for i := l - len(s); i > 0; i-- {
		s = "0" + s
	}
	return s
}
