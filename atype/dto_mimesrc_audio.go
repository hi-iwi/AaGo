package atype

import (
	"github.com/hi-iwi/AaGo/aenum"
	"strings"
)

type AudioSrc struct {
	Processor int    `json:"processor"`
	Pattern   string `json:"pattern"` // e.g.  https://xxx/video.avi?quality=${QUALITY}
	Origin    string `json:"origin"`  // 不一定是真实的
	Path      string `json:"path"`
	//Filename  string `json:"filename"` // basename + extension   直接交path给服务端处理
	Filetype aenum.FileType `json:"filetype"` // aenum.Filetype.Int8()
	Size     Uint24         `json:"size"`     // atype.Uint24.Int8()
	Duration uint16         `json:"duration"` // 时长，秒
}

func (s AudioSrc) Filename() Audio { return NewAudio(s.Path, true) }

func (s AudioSrc) Adjust(quality string) string {
	return strings.ReplaceAll(s.Pattern, "${QUALITY}", quality)
}
