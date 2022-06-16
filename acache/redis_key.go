package acache

import "time"

const (
	HourlyTtl = 24 * time.Hour     // 要求每小时会自动清除之前表；为了避免宕机等影响，ttl设计长一点，24小时内宕机恢复，就能使用
	DailyTtl  = 3 * 24 * time.Hour // 要求每天会自动清除之前表；为了避免宕机等影响，ttl设计长一点；
)

// 使当前时段的放到最后
func BatchKeys(n uint8, final uint8, ignoreCurrent bool, f func(uint8) string) []string {
	l := final
	if ignoreCurrent {
		l--
	}
	ks := make([]string, l)
	var i, j uint8
	// 当前时段一定要在最后
	if n+1 < final {
		for i = n + 1; i < final; i++ {
			ks[j] = f(i)
			j++
		}
	}

	for i = 0; i < n; i++ {
		ks[j] = f(i)
		j++
	}
	if ignoreCurrent {
		return ks
	}

	ks[j] = f(n)

	return ks
}

// final 不包括
func Key(n uint8, final uint8, prev bool, f func(uint8) string) string {
	if prev {
		n--
	}
	if n < 0 || n >= final {
		n = final - 1
	}
	return f(n)
}

func hourIdx(interval uint8) (uint8, uint8) {
	max := uint8(24) // 一天24小时，因数：1，2，3，4，6，8，12，24
	if interval > (max/2) || interval == 0 {
		return 1, 1
	}
	if interval > max {
		interval = max
	}
	final := max / interval
	n := uint8(time.Now().Hour()) % final
	return n, final
}

// @param interval uint8  一天内，每interval小时一张表，最大不能超过24，且被24整除，1，2，3，4，6，8，12，24
func HourlyKey(interval uint8, prev bool, f func(uint8) string) string {
	n, final := hourIdx(interval)
	return Key(n, final, prev, f)
}
func HourlyKeys(interval uint8, ignoreCurrent bool, f func(uint8) string) []string {
	n, final := hourIdx(interval)
	return BatchKeys(n, final, ignoreCurrent, f)
}

func dayIdx(interval uint8) (uint8, uint8) {
	max := uint8(30) // 30 因数：1 2 3 5 6 10 15 30  因数越多，能够拆分的可能性越多，所以选择每个月30天
	if interval > (max/2) || interval == 0 {
		return 1, 1
	}
	if interval > max {
		interval = max
	}
	final := max / interval
	n := uint8(time.Now().Day()) % final
	return n, final
}

// @param interval uint8 一个月内，每interval天一张表
func DailyKey(interval uint8, prev bool, f func(uint8) string) string {
	n, final := dayIdx(interval)
	return Key(n, final, prev, f)
}

func DailyKeys(interval uint8, ignoreCurrent bool, f func(uint8) string) []string {
	n, final := dayIdx(interval)
	return BatchKeys(n, final, ignoreCurrent, f)
}
