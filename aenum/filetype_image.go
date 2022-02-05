package aenum

import (
	"strconv"
)

type ImageType uint16

const (
	UnknownImageType ImageType = 0
	Jpeg             ImageType = 1
	Png              ImageType = 2
	Gif              ImageType = 3
	Webp             ImageType = 4
)

func NewImageType(mime string) (ImageType, bool) {
	switch mime {
	case "jpg", ".jpg", "jpeg", ".jpeg", "image/jpeg":
		return Jpeg, true
	case "png", ".png", "image/png":
		return Png, true
	case "gif", ".gif", "image/gif":
		return Gif, true
	case "webp", ".webp", "image/webp":
		return Webp, true
	}
	return UnknownImageType, false
}
func (t ImageType) Valid() bool {
	return t > UnknownImageType && t <= Webp
}

func (t ImageType) Raw() uint16 {
	return uint16(t)
}
func (t ImageType) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

func (t ImageType) Name() string {
	switch t {
	case Jpeg:
		return "jpg"
	case Png:
		return "png"
	case Gif:
		return "gif"
	case Webp:
		return "webp"
	}
	return t.String()
}

func (t ImageType) Ext() string {
	return "." + t.Name()
}

func (t ImageType) ContentType() string {
	switch t {
	case Jpeg:
		return "image/jpeg"
	case Png:
		return "image/png"
	case Gif:
		return "image/gif"
	case Webp:
		return "image/webp"
	}
	return ""
}
func (t ImageType) Is(t2 ImageType) bool {
	return t == t2
}
func (t ImageType) In(ts []ImageType) bool {
	for _, ty := range ts {
		if ty == t {
			return true
		}
	}
	return false
}
