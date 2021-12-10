package adto

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strconv"
)

// 存储在数据库里面，图片列表
type ImgSrc struct {
	//Url    string `json:"url"`
	Path   string `json:"path"`
	Width  uint16 `json:"width"`
	Height uint16 `json:"height"`
}
type VideoSrc struct {
	//Url  string `json:"url"`
	Path     string `json:"path"`
	Width    uint16 `json:"width"`
	Height   uint16 `json:"height"`
	Duration uint32 `json:"duration"` // 时长，秒
	Size     uint32 `json:"size"`
}



func ParseImgSrc(content string) string {
	reg, _ := regexp.Compile(`<img([^>]+)data-path="([^"]+)"([^>]*)>`)
	dataReg, _ := regexp.Compile(`data-([a-z]+)="([^"]+)"`)
	matches := reg.FindAllStringSubmatch(content, -1)
	imgs := make([]ImgSrc, len(matches))
	for i, match := range matches {
		ni := ImgSrc{
			Path: match[2],
		}
		data := match[1] + match[3]

		dataMatches := dataReg.FindAllStringSubmatch(data, -1)
		for _, dm := range dataMatches {
			switch dm[1] {
			case "width":
				w, _ := strconv.Atoi(dm[2])
				ni.Width = uint16(w)
			case "height":
				h, _ := strconv.Atoi(dm[2])
				ni.Height = uint16(h)
			}
		}
		imgs[i] = ni
	}

	buf := bytes.NewBuffer([]byte{})
	je := json.NewEncoder(buf)
	je.SetEscapeHTML(false)
	je.Encode(imgs)

	return buf.String()
}

type ImagePattern struct {
	Height      int
	Width       int
	Quality     int
	MaxWidth    int
	MaxHeight   int
	WatermarkID int
}

func NewImagePattern(tag ...string) *ImagePattern {
	p := &ImagePattern{}
	if len(tag) > 0 {
		p.Parse(tag[0])
	}
	return p
}

func (p *ImagePattern) Parse(tag string) {
	reg, _ := regexp.Compile(`([a-z]+)(\d+)`)
	matches := reg.FindAllStringSubmatch(tag, -1)
	for _, match := range matches {
		v, _ := strconv.Atoi(match[2])
		/**
		 * w width, h height, q quanlity, v max width, g max height
		 *    	img.width <= v ,   img.width = w  两者区别
		 * xN  有意义，对于不定尺寸的白名单，自动化方案是：先获取 x1 的尺寸，然后 xN ，之后把 source 裁剪
		 */
		t := match[1]
		switch t {
		case "h":
			p.Height = v
		case "w":
			p.Width = v
		case "g":
			p.MaxHeight = v
		case "v":
			p.MaxWidth = v
		case "q":
			p.Quality = v
		case "m":
			p.WatermarkID = v
		}
	}
}


