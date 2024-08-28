package atype

import (
	"github.com/hi-iwi/AaGo/aenum"
)

type DocumentSrc struct {
	Provider int            `json:"provider"`
	Path     string         `json:"path"`
	Url      string         `json:"url"`
	Filetype aenum.FileType `json:"filetype"` // aenum.Filetype.Int8()
	Size     int            `json:"size"`     // atype.Uint24.Int8()
	Jsonkey  string         `json:"jsonkey"`  // 特殊约定字段
}
