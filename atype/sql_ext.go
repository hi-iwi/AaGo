package atype

import (
	"encoding/json"
	"strings"
)

type File string
type Image string // varchar(55)   45 + 5(.webp) + 5 扩展
type Video string // varchar(55)
type Audio string // varchar(55)
type Files struct{ NullStrings }
type Images struct{ NullStrings }
type Videos struct{ NullStrings }
type Audios struct{ NullStrings }

// 仅保留文件名，去掉目录
func trimDir(p string) string {
	if p == "" {
		return ""
	}
	i := strings.LastIndexByte(p, '/')
	if i == len(p) {
		return ""
	}
	return p[i+1:]
}
func NewFile(p string, filenameOnly bool) File {
	if filenameOnly {
		p = trimDir(p)
	}
	return File(p)
}
func MakeFile(img *FileSrc, filenameOnly bool) File {
	if img == nil {
		return ""
	}
	return NewFile(img.Path, filenameOnly)
}
func NewImage(p string, filenameOnly bool) Image {
	if filenameOnly {
		p = trimDir(p)
	}
	return Image(p)
}
func MakeImage(img *ImgSrc, filenameOnly bool) Image {
	if img == nil {
		return ""
	}
	return NewImage(img.Path, filenameOnly)
}
func NewVideo(p string, filenameOnly bool) Video {
	if filenameOnly {
		p = trimDir(p)
	}
	return Video(p)
}
func MakeVideo(video *VideoSrc, filenameOnly bool) Video {
	if video == nil {
		return ""
	}
	return NewVideo(video.Path, filenameOnly)
}
func NewAudio(p string, filenameOnly bool) Audio {
	if filenameOnly {
		p = trimDir(p)
	}
	return Audio(p)
}
func MakeAudio(audio *VideoSrc, filenameOnly bool) Audio {
	if audio == nil {
		return ""
	}
	return NewAudio(audio.Path, filenameOnly)
}
func (p File) String() string                               { return string(p) }
func (p File) Src(filler func(string) *FileSrc) *FileSrc    { return filler(p.String()) }
func (p Image) String() string                              { return string(p) }
func (p Image) Src(filler func(string) *ImgSrc) *ImgSrc     { return filler(p.String()) }
func (p Video) String() string                              { return string(p) }
func (p Video) Src(filler func(string) *VideoSrc) *VideoSrc { return filler(p.String()) }
func (p Audio) String() string                              { return string(p) }
func (p Audio) Src(filler func(string) *AudioSrc) *AudioSrc { return filler(p.String()) }

func NewFiles(s string) Files {
	var x Files
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToFiles(v []string, filenameOnly bool) Files {
	if len(v) == 0 {
		return Files{}
	}
	if filenameOnly {
		for i, s := range v {
			v[i] = trimDir(s)
		}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Files{}
	}

	return NewFiles(string(s))
}

func (im Files) Srcs(filler func(path string) *FileSrc) []FileSrc {
	if !im.Valid || im.String == "" {
		return nil
	}
	ims := im.Strings()
	srcs := make([]FileSrc, 0, len(ims))
	for _, im := range ims {
		if im != "" {
			if fi := filler(im); fi != nil {
				srcs = append(srcs, *filler(im))
			}
		}
	}
	return srcs
}

func NewImages(s string) Images {
	var x Images
	if s != "" {
		x.Scan(s)
	}
	return x
}
func ToImages(v []Image) Images {
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Images{}
	}

	return NewImages(string(s))
}
func ToImages2(v []string, filenameOnly bool) Images {
	if len(v) == 0 {
		return Images{}
	}
	if filenameOnly {
		for i, s := range v {
			v[i] = trimDir(s)
		}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Images{}
	}

	return NewImages(string(s))
}
func ToImages3(v []ImgSrc, filenameOnly bool) Images {
	if len(v) == 0 {
		return Images{}
	}
	imgs := make([]string, len(v))

	for i, s := range v {
		if filenameOnly {
			imgs[i] = trimDir(s.Path)
		} else {
			imgs[i] = s.Path
		}
	}

	s, _ := json.Marshal(imgs)
	if len(s) == 0 {
		return Images{}
	}

	return NewImages(string(s))
}
func (im Images) Srcs(filler func(path string) *ImgSrc) []ImgSrc {
	if !im.Valid || im.String == "" {
		return nil
	}
	ims := im.Strings()
	srcs := make([]ImgSrc, 0, len(ims))
	for _, im := range ims {
		if im != "" {
			if fi := filler(im); fi != nil {
				srcs = append(srcs, *filler(im))
			}
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
func ToVideos(v []string, filenameOnly bool) Videos {
	if len(v) == 0 {
		return Videos{}
	}
	if filenameOnly {
		for i, s := range v {
			v[i] = trimDir(s)
		}
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
func ToAudios(v []string, filenameOnly bool) Audios {
	if len(v) == 0 {
		return Audios{}
	}
	if filenameOnly {
		for i, s := range v {
			v[i] = trimDir(s)
		}
	}
	s, _ := json.Marshal(v)
	if len(s) == 0 {
		return Audios{}
	}

	return NewAudios(string(s))
}

func (im Files) Files(filenameOnly bool) []File {
	imgs := im.Strings()
	if len(imgs) == 0 {
		return nil
	}
	ims := make([]File, len(imgs))
	for i, img := range imgs {
		ims[i] = NewFile(img, filenameOnly)
	}
	return ims
}
func (im Images) Images(filenameOnly bool) []Image {
	imgs := im.Strings()
	if len(imgs) == 0 {
		return nil
	}
	ims := make([]Image, len(imgs))
	for i, img := range imgs {
		ims[i] = NewImage(img, filenameOnly)
	}
	return ims
}
func (im Videos) Videos(filenameOnly bool) []Video {
	imgs := im.Strings()
	if len(imgs) == 0 {
		return nil
	}
	ims := make([]Video, len(imgs))
	for i, img := range imgs {
		ims[i] = NewVideo(img, filenameOnly)
	}
	return ims
}
func (im Audios) Audios(filenameOnly bool) []Audio {
	imgs := im.Strings()
	if len(imgs) == 0 {
		return nil
	}
	ims := make([]Audio, len(imgs))
	for i, img := range imgs {
		ims[i] = NewAudio(img, filenameOnly)
	}
	return ims
}
