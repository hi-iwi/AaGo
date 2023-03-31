package afmt_test

import (
	"github.com/hi-iwi/AaGo/afmt"
	"testing"
	"time"
)

func TestTimeDiff(t *testing.T) {
	a, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	b, _ := time.Parse("2006-01-02 15:04:05", "2022-11-19 20:48:08")
	d1 := afmt.TimeDiff("（%Y年%M个月%D天%H个小时%I分钟%S秒%%）", a, b)
	d2 := "（16年10个月17天5个小时44分钟3秒）"
	if d1 != d2 {
		t.Errorf("afmt.TimeDiff failed " + d1)
	}
	a, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	b, _ = time.Parse("2006-01-02 15:04:05", "2022-11-02 15:04:05")
	d1 = afmt.TimeDiff("（%Y年%M个月%D天%H个小时%I分钟%S秒%%）", a, b)
	d2 = "（16年10个月）"
	if d1 != d2 {
		t.Errorf("afmt.TimeDiff failed " + d1)
	}
	a, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	b, _ = time.Parse("2006-01-02 15:04:05", "2022-01-05 15:04:05")
	d1 = afmt.TimeDiff("（%Y年%M个月%D天%H个小时%I分钟%S秒%%）", a, b)
	d2 = "（16年3天）"
	if d1 != d2 {
		t.Errorf("afmt.TimeDiff failed " + d1)
	}

	a, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	b, _ = time.Parse("2006-01-02 15:04:05", "2022-01-05 15:04:05")
	d1 = afmt.TimeDiff("（%Y年%M个月%D天%H个小时%I分钟%S秒%%）", a, b)
	d2 = "（16年3天）"
	if d1 != d2 {
		t.Errorf("afmt.TimeDiff failed " + d1)
	}
}

func TestDurationString(t *testing.T) {
	a := 49*time.Hour + 1*time.Minute + 0*time.Second
	d, _ := afmt.DurationString(a, "天`小时`分`秒")
	if "2天1小时1分0秒" != d {
		t.Errorf("duration string (%s) should not be %s", a.String(), d)
	}
	d, _ = afmt.DurationString(a, " Days And ` Hours ` Minutes ` Duration", " Day ` Hour ` Minute ` Second")
	if "2 Days And 1 Hour 1 Minute 0 Second" != d {
		t.Errorf("duration string (%s) should not be %s", a.String(), d)
	}
}
