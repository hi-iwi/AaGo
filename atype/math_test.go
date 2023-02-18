package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

func TestMoney(t *testing.T) {
	m := atype.Umoney(1888000)
	if m.Fmt() != "188.8" {
		t.Errorf("atype.Umoney(1888000).Fmt() : %s != 188.8", m.Fmt())
	}
	if m.Format() != "188.80" {
		t.Errorf("atype.Umoney(1888000).Fmt() : %s != 188.80", m.Fmt())
	}
}
func TestPercent(t *testing.T) {
	p, _ := atype.NewPercent(8 * int16(atype.PercentMultiplier))
	if p.Value() != 0.08 {
		t.Errorf("atype.NewPercent(8%%) : %f != 0.08", p.Value())
	}

}
