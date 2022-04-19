package adto

import (
	"strconv"
	"strings"
)

// 存储在数据库里面，图片列表，为了节省空间，用数组来
// 数据库存储方式为 dtype.NullImgSrc，即  [path,size,width,height]
type ImgSrc struct {
	Processor int    `json:"processor"` // 图片处理ID，如阿里云图片处理、网易云图片处理等
	Fill      string `json:"fill"`      // e.g.  https://xxx/img.jpg?width=${WIDTH}&height=${HEIGHT}
	Fit       string `json:"fit"`       // e.g. https://xxx/img.jpg?maxwidth=${MAXWIDTH}
	Path      string `json:"path"`      // path 可能是 filename，也可能是 带文件夹的文件名
	Filetype  uint16 `json:"filetype"`  // aenum.Filetype.Raw()
	Size      uint32 `json:"size"`      // dtype.Uint24.Raw()
	Width     uint16 `json:"width"`
	Height    uint16 `json:"height"`
}

type VideoSrc struct {
	Processor int    `json:"processor"`
	Fit       string `json:"fit"` // e.g.  https://xxx/video.avi?quality=${QUALITY}
	Path      string `json:"path"`
	Filetype  uint16 `json:"filetype"` // aenum.Filetype.Raw()
	Size      uint32 `json:"size"`     // dtype.Uint24.Raw()
	Width     uint16 `json:"width"`
	Height    uint16 `json:"height"`
	Duration  uint32 `json:"duration"` // 时长，秒
}

func (s ImgSrc) FillTo(width, height uint16) string {
	u := s.Fill
	u = strings.ReplaceAll(u, "${WIDTH}", strconv.FormatUint(uint64(width), 10))
	u = strings.ReplaceAll(u, "${HEIGHT}", strconv.FormatUint(uint64(height), 10))
	return u
}

func (s ImgSrc) FitTo(maxwidth uint16) string {
	u := s.Fit
	return strings.ReplaceAll(u, "${MAXWIDTH}", strconv.FormatUint(uint64(maxwidth), 10))
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
	for i, path := range paths {
		srcs[i] = filler(path)
	}
	return srcs
}
