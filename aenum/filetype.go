package aenum

import (
	"strconv"
	"strings"
)

// filetype 统一了，方便客户端分析 path/filetype 结构类型
type FileType uint16

const (
	UnknownType FileType = 0

	// 图片类型范围：1-9999
	Jpeg       FileType = 1
	Png        FileType = 2
	Gif        FileType = 3
	Webp       FileType = 4
	Heic       FileType = 5 // iPhone 拍摄的照片
	OtherImage FileType = 9999

	// 音频类型范围：10000-19999
	Mp3        FileType = 10000
	X3pg       FileType = 10001
	X3pg2      FileType = 10002
	Aiff       FileType = 10003
	OtherAudio FileType = 19999

	// 视频范围：20000-29999
	Avi        FileType = 20000
	Mov        FileType = 20001 // Apple QuickTime
	Mpeg       FileType = 20002
	Mp4        FileType = 20003 // MPEG-4
	X3gp       FileType = 20004
	X3gp2      FileType = 20005
	Webm       FileType = 20006
	Wav        FileType = 20007
	OtherVideo FileType = 29999

	// 文件范围
)

func NewImageType(mime string) (FileType, bool) {
	switch strings.ToLower(mime) {
	case "jpg", ".jpg", "jpeg", ".jpeg", "image/jpeg":
		return Jpeg, true
	case "png", ".png", "image/png":
		return Png, true
	case "gif", ".gif", "image/gif":
		return Gif, true
	case "webp", ".webp", "image/webp":
		return Webp, true
	case "heic", ".heic", "image/heic", "heif", ".heif", "image/heif":
		return Heic, true
	}
	return OtherImage, false
}

func NewAudioType(mime string) (FileType, bool) {
	switch mime {
	case "mp3", ".mp3", "audio/mpeg":
		return Mp3, true
	case "3gp", ".3gp", "audio/3gpp":
		return X3pg, true
	case "3g2", ".3g2", "audio/3gpp2":
		return X3pg2, true
	case "aiff", ".aiff", "aif", ".aif", "aifc", ".aifc", "audio/x-aiff":
		return Aiff, true
	}
	return OtherAudio, false
}

func NewVideoType(mime string) (FileType, bool) {
	switch mime {
	case "avi", ".avi", "video/x-msvideo":
		return Avi, true
	case "mov", ".mov", "video/quicktime":
		return Mov, true
	case "mpeg", ".mpeg", "video/mpeg":
		return Mpeg, true
	case "mp4", ".mp4", "video/mp4":
		return Mp4, true
	case "3gp", ".3gp", "video/3gpp":
		return X3gp, true
	case "3g2", ".3g2", "video/3gpp2":
		return X3gp2, true
	case "webm", ".webm", "video/webm":
		return Webm, true
	case "wav", ".wav", "video/x-wav":
		return Wav, true
	}
	return OtherVideo, false
}

func (t FileType) Uint16() uint16    { return uint16(t) }
func (t FileType) String() string    { return strconv.FormatUint(uint64(t), 10) }
func (t FileType) Is(t2 uint16) bool { return t.Uint16() == t2 }
func (t FileType) In(ts ...FileType) bool {
	for _, ty := range ts {
		if ty == t {
			return true
		}
	}
	return false
}

func (t FileType) Name() string {
	switch t {
	case Jpeg:
		return "jpg"
	case Png:
		return "png"
	case Gif:
		return "gif"
	case Webp:
		return "webp"
	case Heic:
		return "heic"

	case Mp3:
		return "mp3"
	case X3pg:
		return "3gp"
	case X3pg2:
		return "3g2"
	case Aiff:
		return "aiff"
	case Avi:
		return "avi"
	case Mov:
		return "mov"
	case Mpeg:
		return "mpeg"
	case Mp4:
		return "mp4"
	//case X3gp:
	//	return "3gp"
	//case X3gp2:
	//	return "3g2"
	case Webm:
		return "webm"
	case Wav:
		return "wav"
	}
	return "unknown" // 默认是 jpeg，跟 jpg 区分开来
}

func (t FileType) Ext() string { return "." + t.Name() }

func (t FileType) ContentType() string {
	switch t {
	case Jpeg:
		return "image/jpeg"
	case Png:
		return "image/png"
	case Gif:
		return "image/gif"
	case Webp:
		return "image/webp"
	case Heic:
		return "image/heic"

	}
	return ""
}
