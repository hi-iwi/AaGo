package aenum

import "strconv"

type VideoType uint16

const (
	UnknownVideoType VideoType = 0
	Avi              VideoType = 1
	Mov              VideoType = 2 // Apple QuickTime
	Mpeg             VideoType = 3
	Mp4              VideoType = 4 // MPEG-4
	X3gp             VideoType = 5
	X3gp2            VideoType = 6
	Webm             VideoType = 7
	Wav              VideoType = 8
)

func NewVideoType(mime string) (VideoType, bool) {
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
	return UnknownVideoType, false
}
func (t VideoType) Uint16() uint16    { return uint16(t) }
func (t VideoType) String() string    { return strconv.FormatUint(uint64(t), 10) }
func (t VideoType) Is(t2 uint16) bool { return t.Uint16() == t2 }
func (t VideoType) In(ts ...VideoType) bool {
	for _, ty := range ts {
		if ty == t {
			return true
		}
	}
	return false
}

func (t VideoType) Name() string {
	switch t {
	case Avi:
		return "avi"
	case Mov:
		return "mov"
	case Mpeg:
		return "mpeg"
	case Mp4:
		return "mp4"
	case X3gp:
		return "3gp"
	case X3gp2:
		return "3g2"
	case Webm:
		return "webm"
	case Wav:
		return "wav"
	}
	return t.String()
}
func (t VideoType) Ext() string { return "." + t.Name() }
