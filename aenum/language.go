package aenum

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
