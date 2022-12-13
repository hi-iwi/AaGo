package atype

import (
	"github.com/hi-iwi/AaGo/aenum"
)

type VideoSrc struct {
	Processor int    `json:"processor"`
	Fit       string `json:"fit"`    // e.g.  https://xxx/video.avi?quality=${QUALITY}
	Origin    string `json:"origin"` // 不一定是真实的
	Path      string `json:"path"`
	//Filename  string `json:"filename"` // basename + extension   直接交path给服务端处理
	Filetype aenum.VideoType `json:"filetype"` // aenum.Filetype.Int8()
	Size     Uint24          `json:"size"`     // atype.Uint24.Int8()
	Width    uint16          `json:"width"`
	Height   uint16          `json:"height"`
	Duration uint16          `json:"duration"` // 时长，秒
	Allowed  [][2]uint16     `json:"allowed"`  // 限定允许的width,height
}

func (s VideoSrc) Filename() Video { return ToVideo(s.Path) }
