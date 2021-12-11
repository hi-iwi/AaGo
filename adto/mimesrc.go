package adto

import (
	"bytes"
	"encoding/json"
	"github.com/hi-iwi/AaGo/dtype"
	"regexp"
	"strconv"
)

// 存储在数据库里面，图片列表，为了节省空间，用数组来
type ImgSrc struct {
	//Url    string `json:"url"`
	Path   string `json:"path"`
	Size   uint32 `json:"size"`
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

func ParseImgAdto(m [4]interface{}) ImgSrc {
	var is ImgSrc
	is.Path, _ = m[0].(string)
	//  If you sent the JSON value through browser then any number you sent that will be the type float64
	is.Size = dtype.New(m[1]).DefaultUint32(0)
	is.Width = dtype.New(m[2]).DefaultUint16(0)
	is.Height = dtype.New(m[3]).DefaultUint16(0)
	return is
}
func EncodeImgSrc(content string) json.RawMessage {
	reg, _ := regexp.Compile(`<img([^>]+)data-path="([^"]+)"([^>]*)>`)
	dataReg, _ := regexp.Compile(`data-([a-z]+)="([^"]+)"`)
	matches := reg.FindAllStringSubmatch(content, -1)
	var b bytes.Buffer
	const firstBracket = 1
	b.WriteByte('[')
	// b.Grow()
	for _, match := range matches {
		var size, width, height string
		path := match[2]
		data := match[1] + match[3]
		dataMatches := dataReg.FindAllStringSubmatch(data, -1)
		for _, dm := range dataMatches {
			switch dm[1] {
			case "size":
				size = dm[2]
			case "width":
				width = dm[2]
			case "height":
				height = dm[2]
			}
		}
		if b.Len() > firstBracket {
			b.WriteByte(',')
		}
		b.WriteByte('[')
		b.WriteByte('"')
		b.WriteString(path) // path
		b.WriteByte('"')
		b.WriteByte(',')
		b.WriteString(size)
		b.WriteByte(',')
		b.WriteString(width)
		b.WriteByte(',')
		b.WriteString(height)
		b.WriteByte(']')
	}
	if b.Len() > firstBracket {
		b.WriteByte(']')
		return b.Bytes()
	}
	return nil
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
