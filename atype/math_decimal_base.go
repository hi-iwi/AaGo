package atype

import "math"

type RoundType uint8

const (
	Round        RoundType = 0 // math.Round
	Floor        RoundType = 1 // math.Floor  -> floor(0.1) =0; floor(-0.1) = -1
	Ceil         RoundType = 2 // math.Ceil
	RoundReverse RoundType = 3 // v > 0 ? round(v) : -round(-v)
	RoundTrim    RoundType = 4 // v > 0 ? floor(v) : ceil(v)    -> roundTrim(0.1)=0; roundTrim(-0.1) = 0
	RoundAway    RoundType = 5 // round away from the origin point v > 0 ? math.ceil(v) : math.floor(v)
)

func (r RoundType) IsRound() bool {
	return r == Round
}
func (r RoundType) IsFloor() bool {
	return r == Floor
}
func (r RoundType) IsCeil() bool {
	return r == Ceil
}
func (r RoundType) IsReverse() bool {
	return r == RoundReverse
}
func (r RoundType) IsTrim() bool {
	return r == RoundTrim
}
func (r RoundType) IsAway() bool {
	return r == RoundAway
}

func (r RoundType) Round(x float64) float64 {
	switch r {
	case Floor:
		return math.Floor(x)
	case Ceil:
		return math.Ceil(x)
	case RoundReverse:
		if x > 0 {
			return math.Round(x)
		}
		return -math.Round(-x)
	case RoundTrim:
		if x > 0 {
			return math.Floor(x)
		}
		return math.Ceil(x)
	case RoundAway:
		if x > 0 {
			return math.Ceil(x)
		}
		return math.Floor(x)
	}
	return math.Round(x)
}

// separator:[计] 分隔符,倾向于可显示字符，通常单个字符或字符串
// delimiter:[计] 定界符,用于可显示字符和不可显示字符；通常成对出现，比如（）、《》、""、“” 等
type DecimalFormat struct {
	SegmentSize uint8     // 整数部分每segmentSize位使用separator隔开
	Separator   string    // 整数部分分隔符，如英文每3位一个逗号；中文每4位一个空格等表示方法
	Scale       uint8     // 保留小数位数，0则表示不限制
	TrimScale   bool      // 是否删除小数尾部无效的0
	ScaleRound  RoundType //   @warn 如果进位到整数，则只保留.999...；负数按正数部分round
}

func NewDecimalFormat(format *DecimalFormat) *DecimalFormat {
	if format == nil {
		format = &DecimalFormat{
			SegmentSize: 0,
			Separator:   "",
			Scale:       0,
			TrimScale:   false,
			ScaleRound:  Floor,
		}
	} else {
		if format.SegmentSize > 0 && format.Separator == "" {
			format.Separator = ","
		}
		if format.Scale > DecimalScale {
			format.Scale = DecimalScale
		}

	}
	return format
}
func formatWhole(s string, segmentSize uint8, separator string) string {
	if separator == "" {
		separator = ",l"
	}
	if segmentSize == 0 || len(s) <= int(segmentSize) {
		return s
	}
	var s2 string
	j := 0
	for i := len(s) - 1; i > -1; i-- {
		if j > 0 && j%int(segmentSize) == 0 {
			s2 = separator + s2
		}
		s2 = string(s[i]) + s2
		j++
	}
	return s2
}

func padRight(str string, pad string, minlen int) string {
	if pad == "" || len(str) >= minlen {
		return str
	}
	for {
		str += pad
		if len(str) >= minlen {
			return str
		}
	}
}
