package atype

// 文本不存在太长的；不要使用 null string，否则插入空字符串比较麻烦
type Text string // Text 65535 bytes
//type MediumText struct{ sql.NullString } // MediumText 16777215 bytes
//type LongText struct{ sql.NullString }   // LongText 4294967295 bytes

// HTML 一律采用 template.HTML
//type Html string       // TEXT保存的 HTML 格式
//type MediumHtml string // MediumText 16777215 bytes
//type LongHtml string   // LongText 4294967295 bytes
