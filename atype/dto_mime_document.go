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
	Checksum string         `json:"checksum"` // 图片、视频、音频会被压缩，checksum 无意义；这类文件不能被压缩
	Info     string         `json:"info"`     // 冗余数据
	Jsonkey  string         `json:"jsonkey"`  // 特殊约定字段
}
