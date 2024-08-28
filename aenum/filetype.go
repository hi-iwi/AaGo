package aenum

import (
	"strconv"
)

// filetype 统一了，方便客户端分析 path/filetype 结构类型。也方便客户端上传的格式符合标准格式。
// .3pg 既是音频文件，也是视频文件。因此，不能单纯通过后缀知晓文件类型。需要客户端上传的时候预先知道是音频或视频。
type FileType uint16

func ParseFileType(mime string, types map[FileType][]string) (FileType, bool) {
	if mime == "" {
		return 0, false
	}
	for ft, mimes := range types {
		for _, m := range mimes {
			if m == mime {
				return ft, true
			}
		}
	}
	return 0, false
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
