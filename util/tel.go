package util

import (
	"regexp"
	"strings"
)

func FormatTel(tel string) (string, bool) {
	// (0755)12345678
	tel = strings.ReplaceAll(tel, "(", "")
	tel = strings.ReplaceAll(tel, "（", "")
	tel = strings.ReplaceAll(tel, "﹝", "")
	tel = strings.ReplaceAll(tel, "「", "")

	tel = strings.ReplaceAll(tel, ")", "-")
	tel = strings.ReplaceAll(tel, "）", "-")
	tel = strings.ReplaceAll(tel, "﹞", "-")
	tel = strings.ReplaceAll(tel, "」", "-")
	tel = strings.ReplaceAll(tel, " ", "-")
	// 1xx    0755-12345678
	tel = strings.ReplaceAll(tel, "--", "-")
	tel = strings.ReplaceAll(tel, " ", "")

	// 加拿大手机号7位数字
	matched, _ := regexp.MatchString(`^([1-9]\d{6,}|\d{2,}-[1-9]\d{5,})$`, tel)
	return tel, matched
}
