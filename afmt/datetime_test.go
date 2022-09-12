package afmt_test

import (
	"github.com/hi-iwi/AaGo/afmt"
	"testing"
	"time"
)

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
