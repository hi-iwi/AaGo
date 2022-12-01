package aenum

import "strconv"

type Language uint8

const (
	NilLanguage Language = 0
	EnUS        Language = 1
	ZhCN        Language = 86 // simpified chinese

)

var LanguageCodes = map[Language]string{
	EnUS: "en-US",
	ZhCN: "zh-CN",
}
var LanguageNames = map[Language]map[Language]string{
	ZhCN: {
		EnUS: "英语",
		ZhCN: "中文",
	},
	EnUS: {
		EnUS: "English",
		ZhCN: "Chinese",
	},
}

func NewLanguage(lang uint8) (Language, bool) {
	return Language(lang), true
}
func (c Language) Uint8() uint8    { return uint8(c) }
func (c Language) String() string  { return strconv.FormatUint(uint64(c), 10) }
func (c Language) Is(x uint8) bool { return c.Uint8() == x }
func (c Language) In(args ...Language) bool {
	for _, a := range args {
		if a == c {
			return true
		}
	}
	return false
}
