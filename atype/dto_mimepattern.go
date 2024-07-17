package atype

import (
	"regexp"
	"strconv"
)

type VideoPattern struct {
}
type ImagePattern struct {
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	Quality   uint8  `json:"quality"`
	MaxWidth  int    `json:"max_width"`
	MaxHeight int    `json:"max_height"`
	Watermark string `json:"watermark"`
}

func ImageFill(width, height int) ImagePattern {
	return ImagePattern{Width: width, Height: height}
}
func ImageFitWidth(maxWidth int) ImagePattern {
	return ImagePattern{MaxWidth: maxWidth}
}
func ToImagePattern(tag string) ImagePattern {
	reg, _ := regexp.Compile(`([a-z]+)(\d+)`)
	matches := reg.FindAllStringSubmatch(tag, -1)
	var p ImagePattern
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
			p.Quality = uint8(v)
		case "k":
			p.Watermark = match[2]
		}
	}
	return p
}

func ToVideoPattern(tag string) VideoPattern {
	return VideoPattern{}
}
