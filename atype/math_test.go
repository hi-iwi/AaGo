package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

func TestMoney(t *testing.T) {
	m := atype.Umoney(188.8 * float64(atype.Yuan))
	if m.Fmt() != "188.8" {
		t.Errorf("atype.Umoney(1888000).Fmt() : %s != 188.8", m.Fmt())
	}
	if m.Format() != "188.80" {
		t.Errorf("atype.Umoney(1888000).Fmt() : %s != 188.80", m.Fmt())
	}
	p := atype.ToPercent(0.7)
	a := atype.NewYuan(2360)
	s := a.MulPercent(p).Fmt()
	if s != "16.52" {
		t.Errorf("atype.Umoney(1888000).Fmt() : %s != 188.80", m.Fmt())
	}
}
func TestPercent(t *testing.T) {
	p := atype.ToPercent(0.7)
	if p.Value().String() != "0.007" {
		t.Errorf("atype.NewPercent(0.7%%) : %s != 0.007", p.Value().String())
	}
	if p.Percent().String() != "0.7" {
		t.Errorf("atype.NewPercent(0.7%%) : %s != 0.7", p.Percent().String())
	}
}
