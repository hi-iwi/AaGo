package atype

import (
	"github.com/hi-iwi/AaGo/aenum"
	"github.com/hi-iwi/AaGo/util"
	"path"
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

// 绝对asset url，不需要进行任何图片处理
func NewImgAsset(url string) *ImgSrc {
	if url == "" {
		return nil
	}
	ft, _ := aenum.NewImageType(path.Ext(url))
	return &ImgSrc{
		Processor: 0,
		Fill:      url,
		Fit:       url,
		Origin:    url,
		Path:      url,
		Filetype:  ft,
		Size:      0,
		Width:     0,
		Height:    0,
		Allowed:   nil,
	}
}
 
// similar to path.Base()
func (s ImgSrc) Filename() string { return util.Filename(s.Path) }

func (s ImgSrc) FillTo(width, height uint16) string {
	if s.Processor == 0 {
		return s.Origin
	}
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
	if s.Processor == 0 {
		return s.Origin
	}
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
