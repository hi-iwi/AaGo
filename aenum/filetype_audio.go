package aenum

import "strconv"

type AudioType uint16

const (
	UnknownAudioType AudioType = 0
	Mp3              AudioType = 1
	X3pg             AudioType = 2
	X3pg2            AudioType = 3
)

func NewAudioType(mime string) (AudioType, bool) {
	switch mime {
	case "mp3", ".mp3", "audio/mpeg":
		return Mp3, true
	case "3gp", ".3gp", "audio/3gpp":
		return X3pg, true
	case "3g2", ".3g2", "audio/3gpp2":
		return X3pg2, true

	}
	return UnknownAudioType, false
}
func (t AudioType) Valid() bool {return t > UnknownAudioType && t <= X3pg2}

func (t AudioType) Uint8() uint16 {return uint16(t)}
func (t AudioType) String() string {return strconv.FormatUint(uint64(t), 10)}

func (t AudioType) Name() string {
	switch t {
	case Mp3:
		return "mp3"
	case X3pg:
		return "3gp"
	case X3pg2:
		return "3g2"
	}
	return t.String()
}

func (t AudioType) Ext() string {return "." + t.Name()}
func (t AudioType) Is(t2 AudioType) bool {return t == t2}
func (t AudioType) In(ts []AudioType) bool {
	for _, ty := range ts {
		if ty == t {
			return true
		}
	}
	return false
}
