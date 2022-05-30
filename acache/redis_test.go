package acache_test

import (
	"github.com/hi-iwi/AaGo/acache"
	"testing"
)

func TestHourlyKey(t *testing.T) {
	for i := 0; i < 32; i++ {
		x := i/7
		if x > 3 {
			x = 3  // 当月29号之后都并入上一周期
		}
		t.Log(acache.BatchKeys(x, 4, true))
		t.Log(acache.BatchKeys(x, 4, false))
	}
}

func TestWeeklyKey(t *testing.T) {
	for i := 0; i < 32; i++ {
		x := i/7
		if x > 3 {
			x = 3  // 当月29号之后都并入上一周期
		}
		t.Log(acache.BatchKeys(x, 4, true))
		t.Log(acache.BatchKeys(x, 4, false))
	}
}