package adto

import (
	"bytes"
	"encoding/json"
	"github.com/hi-iwi/AaGo/dtype"
	"regexp"
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
