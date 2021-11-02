package afmt

import (
	"bytes"
	"errors"
	"math"
	"time"
)

// 格式化 Duration
// time.Duration.String()  返回如：895061h51m1.00001s
// 这里对系统功能进行扩展，支持多种语言，但是只支持到整数秒。
// layout -> d:h:m:s ，如  Days`Hours`Minutes`Seconds   Hours`Minutes`Seconds  天`小时`分`秒 小时`分`秒
// singlelar -> Day.Hour.Minute.Second D.H.M.S

func DurationString(d time.Duration, layout string, singlelar ...string) (string, error) {

	var x int
	g := len(singlelar) > 0

	la := bytes.Split([]byte(layout), []byte{'`'})
	if len(la) < 3 {
		return "", errors.New("layout pattern must be like h`m`s or d`h`m`s")
	}
	for _, p := range la {
		x += len(p)
	}
	sa := [][]byte{{'h'}, {'m'}, {'s'}}
	if g {
		sa = bytes.Split([]byte(singlelar[0]), []byte{'`'})
		if len(sa) != len(la) {
			return "", errors.New("the singular layout pattern does not match the plural layout pattern")
		}
		x = 0
		for i, p := range la {
			if len(sa[i]) > len(p) {
				x += len(sa[i])
			} else {
				x += len(p)
			}
		}
	}

	// Largest time is 2540400h10m10s
	l := 22 - 3 + x
	buf := make([]byte, l)
	w := len(buf)
	u := uint64(math.Floor(d.Seconds()))
	neg := d < 0
	if neg {
		u = -u
	}
	n := u % 60

	var pattern []byte
	if g && n < 2 {
		pattern = sa[len(sa)-1] // second's singluar pattern
		w -= len(pattern)
		copy(buf[w:], sa[len(sa)-1])
	} else {
		pattern = la[len(la)-1] // second's pattern
		w -= len(pattern)
		copy(buf[w:], la[len(la)-1]) // end with second's pattern, e.g. "s", "Second", "Seconds", "秒"
	}

	// u is now integer seconds
	w = fmtInt(buf[:w], n)
	u /= 60

	// u is now integer minutes
	if u > 0 {
		n = u % 60
		if g && n < 2 {
			pattern = sa[len(sa)-2] // minute's singluar pattern
			w -= len(pattern)
			copy(buf[w:], sa[len(sa)-2])
		} else {
			pattern = la[len(la)-2] // minute's pattern
			w -= len(pattern)
			copy(buf[w:], la[len(la)-2])
		}
		w = fmtInt(buf[:w], n)
		u /= 60

		// u is now integer hours
		if u > 0 {
			// Better stop at hours because days can be different lengths.
			if len(la) == 3 {
				if g && u < 2 {
					pattern = sa[len(sa)-3] // hour's singluar pattern
					w -= len(pattern)
					copy(buf[w:], sa[len(sa)-3])
				} else {
					pattern = la[len(la)-3] // hour's pattern
					w -= len(pattern)
					copy(buf[w:], la[len(la)-3])
				}

				w = fmtInt(buf[:w], u)
			} else {
				n = u % 24
				if g && n < 2 {
					pattern = sa[len(sa)-3] // hour's singluar pattern
					w -= len(pattern)
					copy(buf[w:], sa[len(sa)-3])
				} else {
					pattern = la[len(la)-3] // hour's pattern
					w -= len(pattern)
					copy(buf[w:], la[len(la)-3])
				}

				w = fmtInt(buf[:w], n)

				u /= 24
				if u > 0 {
					if g && u < 2 {
						pattern = sa[len(sa)-4] // day's singluar pattern
						w -= len(pattern)
						copy(buf[w:], sa[len(sa)-4])
					} else {
						pattern = la[len(la)-4] // day's pattern
						w -= len(pattern)
						copy(buf[w:], la[len(la)-4])
					}

					w = fmtInt(buf[:w], u)
				}
			}

		}
	}

	if neg {
		w--
		buf[w] = '-'
	}

	return string(buf[w:]), nil
}

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}
