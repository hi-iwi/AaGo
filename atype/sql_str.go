package atype

import (
	"html/template"
	"regexp"
	"strings"
)

// 文本不存在太长的；不要使用 null string，否则插入空字符串比较麻烦
type Text string // Text 65535 bytes
//type MediumText struct{ sql.NullString } // MediumText 16777215 bytes
//type LongText struct{ sql.NullString }   // LongText 4294967295 bytes

// HTML 一律采用 template.HTML
//type Html string       // TEXT保存的 HTML 格式
//type MediumHtml string // MediumText 16777215 bytes
//type LongHtml string   // LongText 4294967295 bytes

func (t Text) String() string {
	return string(t)
}

func NewText(s string, trim bool) Text {
	if s == "" {
		return ""
	}
	if strings.Index(s, "<br") > 0 {
		s = strings.ReplaceAll(s, "<br>", "\r\n")
		s = strings.ReplaceAll(s, "<br/>", "\r\n")
	}
	if trim {
		re := regexp.MustCompile(`(^[\r\n\s\t]+)|([\r\n\s\t]$)`)
		s = re.ReplaceAllString(s, "")
		re = regexp.MustCompile(`[\s\t]*[\r\n]+[\s\t]*[\r\n]+[\s\t]*`)
		s = re.ReplaceAllString(s, `\r\n`)
	}
	return Text(s)
}

// 编码的时候
func (t Text) Html() template.HTML {
	if t == "" {
		return ""
	}
	s := t.String()
	if strings.IndexAny(s, `\r\n`) > 0 {
		s = strings.ReplaceAll(s, `\r\n`, "<br>")
		s = strings.ReplaceAll(s, `\r`, "<br>")
		s = strings.ReplaceAll(s, `\n`, "<br>")
	}
	return template.HTML(s)
}
