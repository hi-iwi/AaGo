package util

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strconv"
)

type OasBimg struct {
	//URL    string `json:"url"`
	Path   string `json:"path"`
	Width  uint16 `json:"width"`
	Height uint16 `json:"height"`
	Size   uint32 `json:"size"`
}
type OasBvideo struct {
	//URL  string `json:"url"`
	Path   string `json:"path"`
	Width  uint16 `json:"width"`
	Height uint16 `json:"height"`
	Size   uint32 `json:"size"`
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

func (p *ImagePattern) OasURL(url string) string {
	// if len(style) > 0 {
	// 	g := strings.Split(tag, ".")
	// 	return   g[0] + "-" + style[0] + "." + g[1]
	// }
	return ""
}

func (p *ImagePattern) EastChinaURL(url string) string {
	style := ""
	if p.Height > 0 && p.Width > 0 {
		// y 超出部分要裁剪；enlarge 不足拉伸
		style += strconv.Itoa(p.Width) + "y" + strconv.Itoa(p.Height) + "&enlarge=1"
	} else if p.MaxHeight > 0 {
		style += "0x" + strconv.Itoa(p.MaxHeight)
	} else if p.MaxWidth > 0 {
		style += strconv.Itoa(p.MaxWidth) + "x0"
	}
	return url + "?imageView&thumbnail=" + style
}

func OasMatchBimgs(content string) string {
	reg, _ := regexp.Compile(`<img([^>]+)src="[^"]+"([^>]*)>`)
	dataReg, _ := regexp.Compile(`data-([a-z]+)="([^"]+)"`)
	matches := reg.FindAllStringSubmatch(content, -1)
	imgs := make([]OasBimg, len(matches))
	for i, match := range matches {
		ni := OasBimg{}
		data := match[1] + match[2]

		dataMatches := dataReg.FindAllStringSubmatch(data, -1)
		for _, dm := range dataMatches {
			switch dm[1] {
			case "width":
				w, _ := strconv.Atoi(dm[2])
				ni.Width = uint16(w)
			case "height":
				h, _ := strconv.Atoi(dm[2])
				ni.Height = uint16(h)
			case "size":
				size, _ := strconv.Atoi(dm[2])
				ni.Size = uint32(size)
			case "path":
				ni.Path = dm[2]
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

func OasMatchBvideos(content string) string {
	videos := make([]OasBvideo, 0)
	buf := bytes.NewBuffer([]byte{})
	je := json.NewEncoder(buf)
	je.SetEscapeHTML(false)
	je.Encode(videos)
	return buf.String()
}
