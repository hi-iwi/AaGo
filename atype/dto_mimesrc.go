package atype

import (
	"github.com/hi-iwi/AaGo/aenum"
	"github.com/hi-iwi/AaGo/util"
	"strconv"
	"strings"
)

// 存储在数据库里面，图片列表，为了节省空间，用数组来；具体见 atype.NullStrings or string
type ImgSrc struct {
	Processor int    `json:"processor"` // 图片处理ID，如阿里云图片处理、网易云图片处理等
	Fill      string `json:"fill"`      // e.g.  https://xxx/img.jpg?width=${WIDTH}&height=${HEIGHT}
	Fit       string `json:"fit"`       // e.g. https://xxx/img.jpg?maxwidth=${MAXWIDTH}
	Origin    string `json:"origin"`    // 不一定是真实的
	Path      string `json:"path"`      // path 可能是 filename，也可能是 带文件夹的文件名
	/*
	   不要独立出来 filename，一方面太多内容了；另一方面增加业务侧复杂度
	*/
	//Filename  string `json:"filename"`  // basename + extension  直接交path给服务端处理
	Filetype aenum.ImageType `json:"filetype"` // aenum.Filetype.Int8()
	Size     Uint24          `json:"size"`     // atype.Uint24.Int8()
	Width    uint16          `json:"width"`
	Height   uint16          `json:"height"`
	Allowed  [][2]uint16     `json:"allowed"` // 允许的width,height
}
type AudioSrc struct {
	Processor int    `json:"processor"`
	Fit       string `json:"fit"`    // e.g.  https://xxx/video.avi?quality=${QUALITY}
	Origin    string `json:"origin"` // 不一定是真实的
	Path      string `json:"path"`
	//Filename  string `json:"filename"` // basename + extension   直接交path给服务端处理
	Filetype aenum.AudioType `json:"filetype"` // aenum.Filetype.Int8()
	Size     Uint24          `json:"size"`     // atype.Uint24.Int8()
	Duration uint16          `json:"duration"` // 时长，秒
}

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

func (s AudioSrc) FitTo(quality string) string {
	return strings.ReplaceAll(s.Fit, "${QUALITY}", quality)
}
func (s VideoSrc) FitTo(quality string) string {
	return strings.ReplaceAll(s.Fit, "${QUALITY}", quality)
}

// similar to path.Base()
func (s ImgSrc) Filename() string { return util.Filename(s.Path) }

func (s ImgSrc) FillTo(width, height uint16) string {
	if width >= s.Width && height >= s.Height && s.Origin != "" {
		return s.Origin
	}
	if s.Allowed != nil {
		var matched, found bool
		var mw, mh uint16
		w := width
		h := height
		for _, a := range s.Allowed {
			aw := a[0]
			ah := a[1]
			if aw == width && ah == height {
				found = true
				break
			}
			if !matched {
				if aw > mw {
					mw = aw
					mh = ah
				}
				// 首先找到比缩放比例大过需求的
				if aw >= w && ah >= h {
					w = aw
					h = ah
					matched = true
				}
			} else {
				// 后面的都跟第一次匹配的比，找到最小匹配
				if aw >= width && aw <= w && ah >= height && ah <= h {
					w = aw
					h = ah
				}
			}
		}
		if !found {
			if !matched {
				width = mw
				height = mh
			} else {
				width = w
				height = h
			}
		}
	}

	sw := strconv.FormatUint(uint64(width), 10)
	sh := strconv.FormatUint(uint64(height), 10)
	u := s.Fill
	u = strings.ReplaceAll(u, "${WIDTH}", sw)
	u = strings.ReplaceAll(u, "${HEIGHT}", sh)
	return u
}

func (s ImgSrc) FitTo(maxWidth uint16) string {
	if maxWidth >= s.Width && s.Origin != "" {
		return s.Origin
	}

	if s.Allowed != nil {
		var matched, found bool
		var mw uint16
		w := maxWidth
		for _, a := range s.Allowed {
			aw := a[0]
			if aw == maxWidth {
				found = true
				break
			}
			if !matched {
				if aw > mw {
					mw = aw
				}
				// 首先找到比缩放比例大过需求的
				if aw >= w {
					w = aw
					matched = true
				}
			} else {
				// 后面的都跟第一次匹配的比，找到最小匹配
				if aw >= maxWidth && aw <= w {
					w = aw
				}
			}
		}
		if !found {
			if !matched {
				maxWidth = mw
			} else {
				maxWidth = w
			}
		}
	}
	sw := strconv.FormatUint(uint64(maxWidth), 10)
	return strings.ReplaceAll(s.Fit, "${MAXWIDTH}", sw)
}

func ToImgSrcPtr(path string, filler func(path string) ImgSrc) *ImgSrc {
	if path == "" {
		return nil
	}
	src := filler(path)
	return &src
}

func ToImgSrcs(paths []string, filler func(path string) ImgSrc) []ImgSrc {
	if len(paths) == 0 {
		return nil
	}
	srcs := make([]ImgSrc, len(paths))
	for i, p := range paths {
		srcs[i] = filler(p)
	}
	return srcs
}
