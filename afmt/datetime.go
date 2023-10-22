package afmt

import (
	"bytes"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
)

func timeDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = y2 - y1
	month = int(M2 - M1)
	day = d2 - d1
	hour = h2 - h1
	min = m2 - m1
	sec = s2 - s1

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}
	return
}
func carryTimeDiff(p map[rune]*int, ls map[rune]bool) map[rune]*int {
	arr := []rune{'S', 'I', 'H', 'D', 'M', 'Y'}
	var g bool
	for _, a := range arr {
		if ls[a] {
			if g {
				*p[a]++
			}
			break
		}
		if *p[a] > 0 {
			g = true
		}
	}

	if *p['S'] > 59 {
		*p['S'] -= 60
		*p['I']++
	}
	if *p['I'] > 59 {
		*p['I'] -= 60
		*p['H']++
	}
	if *p['H'] > 23 {
		*p['H'] -= 24
		*p['D']++
	}
	if *p['D'] > 30 {
		*p['D'] -= 31
		*p['M']++
	}
	if *p['M'] > 11 {
		*p['M'] -= 12
		*p['Y']++
	}

	return p
}
func loadTimeDiff(layout []rune, p map[rune]*int) map[rune]bool {
	ls := map[rune]bool{
		'Y': false,
		'M': false,
		'D': false,
		'H': false,
		'I': false,
		'S': false,
	}

	n := len(layout)
	start := false
	for i := 0; i < n; i++ {
		c := layout[i]
		if start {
			if c == '}' {
				start = false
			}
			continue
		}
		if c != '{' || i > n-3 || layout[i+1] != '%' {
			continue
		}
		r := layout[i+2]
		if _, ok := p[r]; ok {
			ls[r] = true
			i += 2
			start = true
		}
	}
	return ls
}

// 计算两个日期之差
// @param layout:  %Y %M %D %H %I %S  e.g. `{%Y年}{%M个月}`
// @param noCarry  true 尾数忽略；false 尾数后面>0，就+1
func TimeDiff(layout string, d1 time.Time, d2 time.Time, noCarry bool) string {
	if layout == "" {
		return ""
	}
	y, m, d, h, mi, sec := timeDiff(d1, d2)
	if y == 0 && m == 0 && d == 0 && h == 0 && mi == 0 && sec == 0 {
		return ""
	}
	la := []rune(layout)
	p := map[rune]*int{
		'Y': &y,
		'M': &m,
		'D': &d,
		'H': &h,
		'I': &mi,
		'S': &sec,
	}
	ls := loadTimeDiff(la, p)
	if !noCarry {
		p = carryTimeDiff(p, ls)
	}

	if *p['Y'] > 0 && !ls['Y'] {
		m += *p['Y'] * 12
		*p['Y'] = 0
	}
	if *p['M'] > 0 && !ls['M'] {
		*p['D'] += *p['M'] * 30 // 近似天数
		*p['M'] = 0
	}
	if *p['D'] > 0 && !ls['D'] {
		*p['H'] += *p['D'] * 24
		*p['D'] = 0
	}
	if *p['H'] > 0 && !ls['H'] {
		*p['I'] += *p['H'] * 60
		*p['H'] = 0
	}
	if *p['I'] > 0 && !ls['I'] {
		*p['S'] += *p['I'] * 60
		*p['I'] = 0
	}
	var out strings.Builder
	n := len(la)
	start := false
	ignore := false
	for i := 0; i < n; i++ {
		c := la[i]
		if start {
			if c == '}' {
				start = false
				ignore = false
				continue
			}
			if !ignore {
				out.WriteRune(c)
			}
			continue
		}
		if c != '{' || i > n-3 || la[i+1] != '%' {
			out.WriteRune(c)
			continue
		}
		q, ok := p[la[i+2]]
		if !ok {
			out.WriteRune(c)
			continue
		}
		if *q == 0 {
			ignore = true
		} else {
			out.WriteString(strconv.Itoa(*q))
		}
		i += 2
		start = true
	}

	return out.String()
}

func DurationInChinese(d time.Duration) string {
	s, _ := DurationString(d, "天`小时`分`秒")
	return s
}

// 格式化 Duration
// time.Duration.String()  返回如：895061h51m1.00001s
// 这里对系统功能进行扩展，支持多种语言，但是只支持到整数秒。
// layout -> d:h:m:s ，如  Days`Hours`Minutes`Duration   Hours`Minutes`Duration  天`小时`分`秒 小时`分`秒
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
		copy(buf[w:], la[len(la)-1]) // end with second's pattern, e.g. "s", "Second", "Duration", "秒"
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
