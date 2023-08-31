package atype

import (
	"database/sql"
	"html/template"
)

type Text struct{ sql.NullString }       // Text 65535 bytes
type MediumText struct{ sql.NullString } // MediumText 16777215 bytes
type LongText struct{ sql.NullString }   // LongText 4294967295 bytes

type Html struct{ sql.NullString }       // TEXT保存的 HTML 格式
type MediumHtml struct{ sql.NullString } // MediumText 16777215 bytes
type LongHtml struct{ sql.NullString }   // LongText 4294967295 bytes

func (t Html) Html() template.HTML       { return template.HTML(t.String) }
func (t Html) MediumHtml() template.HTML { return template.HTML(t.String) }
func (t Html) LongHtml() template.HTML   { return template.HTML(t.String) }
