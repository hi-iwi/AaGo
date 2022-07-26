package atype

import "encoding/json"

type Image string // varchar(55)   45 + 5(.webp) + 5 扩展
type Video string // varchar(55)
type Audio string // varchar(55)
type Images struct{ NullStrings }
type Videos struct{ NullStrings }
type Audios struct{ NullStrings }

func (im Image) Src(filler func(path string) ImgSrc) *ImgSrc {
	return ToImgSrcPtr(string(im), filler)
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

func (im Images) Srcs(filler func(path string) ImgSrc) []ImgSrc {
	return ToImgSrcs(im.Strings(), filler)
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
