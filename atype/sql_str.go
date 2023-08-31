package atype

import (
	"database/sql"
	"html/template"
)

// 文本不存在太长的，阅读也不方便。太长的，应使用html
type Text struct{ sql.NullString } // Text 65535 bytes
//type MediumText struct{ sql.NullString } // MediumText 16777215 bytes
//type LongText struct{ sql.NullString }   // LongText 4294967295 bytes

type Html struct{ sql.NullString }       // TEXT保存的 HTML 格式
type MediumHtml struct{ sql.NullString } // MediumText 16777215 bytes
type LongHtml struct{ sql.NullString }   // LongText 4294967295 bytes

func NewText(s string) (t Text) {
	if s != "" {
		t.Scan(s)
	}
	return
}

func NewHtml(s string) (t Html) {
	if s != "" {
		t.Scan(s)
	}
	return
}
func NewMediumHtml(s string) (t MediumHtml) {
	if s != "" {
		t.Scan(s)
	}
	return
}
func NewLongHtml(s string) (t LongHtml) {
	if s != "" {
		t.Scan(s)
	}
	return
}
func (t Html) Html() template.HTML       { return template.HTML(t.String) }
func (t Html) MediumHtml() template.HTML { return template.HTML(t.String) }
func (t Html) LongHtml() template.HTML   { return template.HTML(t.String) }
