package atype

import (
	"encoding/json"
	"strings"
)

type Image string // varchar(55)   45 + 5(.webp) + 5 扩展
type Video string // varchar(55)
type Audio string // varchar(55)
type Images struct{ NullStrings }
type Videos struct{ NullStrings }
type Audios struct{ NullStrings }

func NewImage(p string) Image { return Image(p) }

// 仅保留文件名，去掉目录
func ToImage(p string) Image {
	if p == "" {
		return ""
	}
	i := strings.LastIndexByte(p, '/')
	if i == len(p) {
		return ""
	}
	return Image(p[i+1:])
}
func (im Image) String() string {
	return string(im)
}
func (im Image) Src(filler func(path string) *ImgSrc) *ImgSrc {
	return filler(im.String())
}

func NewImages(s string) Images {
	var x Images
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToImages(v []string) Images {
	if len(v) == 0 {
		return Images{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Images{}
	}

	return NewImages(string(s))
}

func (im Images) Srcs(filler func(path string) *ImgSrc) []ImgSrc {
	ims := im.Strings()
	srcs := make([]ImgSrc, 0, len(ims))
	for _, im := range ims {
		if im != "" {
			srcs = append(srcs, *filler(im))
		}
	}
	return srcs
}

func NewVideos(s string) Videos {
	var x Videos
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToVideos(v []string) Videos {
	if len(v) == 0 {
		return Videos{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Videos{}
	}

	return NewVideos(string(s))
}

func NewAudios(s string) Audios {
	var x Audios
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToAudios(v []string) Audios {
	if len(v) == 0 {
		return Audios{}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Audios{}
	}

	return NewAudios(string(s))
}
