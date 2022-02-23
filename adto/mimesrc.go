package adto

import (
	"regexp"
	"strconv"
)

// 存储在数据库里面，图片列表，为了节省空间，用数组来
// 数据库存储方式为 dtype.NullImgSrc，即  [path,size,width,height]
type ImgSrc struct {
	Processor int    `name:"-" json:"processor"`  // 图片处理ID，如阿里云图片处理、网易云图片处理等
	Url       string `name:"-" json:"url"`
	Path      string `name:"path" json:"path"`
	Size      uint32 `name:"size" json:"size"`
	Width     uint16 `name:"width" json:"width"`
	Height    uint16 `name:"height" json:"height"`
}

type VideoSrc struct {
	Processor int    `name:"-" json:"processor"`
	Url       string `name:"-" json:"url"`
	Path      string `name:"path" json:"path"`
	Size      uint32 `name:"size" json:"size"`
	Width     uint16 `name:"width" json:"width"`
	Height    uint16 `name:"height" json:"height"`
	Duration  uint32 `name:"duration" json:"duration"` // 时长，秒
}

func EncodeImgSrc(content string) []ImgSrc {
	reg, _ := regexp.Compile(`<img([^>]+)data-path="([^"]+)"([^>]*)>`)
	dataReg, _ := regexp.Compile(`data-([a-z]+)="([^"]+)"`)
	matches := reg.FindAllStringSubmatch(content, -1)
	imgs := make([]ImgSrc, len(matches))
	for i, match := range matches {
		var size, width, height uint64
		path := match[2]
		data := match[1] + match[3]
		dataMatches := dataReg.FindAllStringSubmatch(data, -1)
		for _, dm := range dataMatches {
			switch dm[1] {
			case "size":
				size, _ = strconv.ParseUint(dm[2], 10, 32)
			case "width":
				width, _ = strconv.ParseUint(dm[2], 10, 16)
			case "height":
				height, _ = strconv.ParseUint(dm[2], 10, 16)
			}
		}
		imgs[i] = ImgSrc{
			Path:   path,
			Size:   uint32(size),
			Width:  uint16(width),
			Height: uint16(height),
		}
	}
	return imgs
}
