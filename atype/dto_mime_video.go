package atype

import (
	"github.com/hi-iwi/AaGo/aenum"
	"strings"
)

type VideoSrc struct {
	Provider int    `json:"provider"`
	Pattern  string `json:"pattern"` // e.g.  https://xxx/video.avi?quality=${QUALITY}
	Origin   string `json:"origin"`  // 不一定是真实的
	Path     string `json:"path"`
	Preview  string `json:"preview"` // 一般是 gif 格式动图，所以不能缩放，直接url即可
	//Filename  string `json:"filename"` // basename + extension   直接交path给服务端处理
	Filetype aenum.FileType `json:"filetype"` // aenum.Filetype.Int8()
	Size     int            `json:"size"`     // atype.Uint24.Int8()
	Width    int            `json:"width"`
	Height   int            `json:"height"`
	Duration int            `json:"duration"` // 时长，秒
	Allowed  [][2]int       `json:"allowed"`  // 限定允许的width,height
	Jsonkey  string         `json:"jsonkey"`  // 特殊约定字段
}

func (s VideoSrc) Filename() Video { return NewVideo(s.Path, true) }
func (s VideoSrc) Adjust(quality string) string {
	return strings.ReplaceAll(s.Pattern, "${QUALITY}", quality)
}
