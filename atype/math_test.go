package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"runtime"
	"testing"
)

func TestPercent(t *testing.T) {
	m := atype.ToMoney(1000.0)
	if m.MulPercent(0.6) != atype.ToMoney(6) {
		t.Errorf("atype.ToMoney(1000).MulPercent(0.6) != atype.ToMoney(6)")
	}
	if m.MulPercent(30) != atype.ToMoney(300) {
		t.Errorf("atype.ToMoney(1000).MulPercent(30) != atype.ToMoney(300)")
	}
}
func TestMoney(t *testing.T) {
	m := atype.ToMoney(188.8)
	if m.Fmt() != "188.8" {
		t.Errorf("atype.ToMoney(1888000).FmtPercent() : %s != 188.8", m.Fmt())
	}
	if m.Format() != "188.80" {
		t.Errorf("atype.ToMoney(1888000).FmtPercent() : %s != 188.80", m.Format())
	}
	p := atype.ToPercent(0.7)
	s := atype.ToMoney(2360).MulRate(p).Fmt()
	if s != "16.52" {
		t.Errorf("2360*0.7%% : %s != 16.52", s)
	}

	b := atype.NewMoney(234242342340503)

	if b.Format(0) != "23424234234" {
		t.Errorf("money (%d).Format(0) ==> string(%s)", b, b.Format(0))
	}
	if b.Format(1) != "23424234234.0" {
		t.Errorf("money (%d).Format(1) ==> string(%s)", b, b.Format(1))
	}
	if b.Format(2) != "23424234234.05" {
		t.Errorf("money (%d).Format(2) ==> string(%s)", b, b.Format(2))
	}
	if b.Format(3) != "23424234234.050" {
		t.Errorf("money (%d).Format(3) ==> string(%s)", b, b.Format(3))
	}
	if b.Format(4) != "23424234234.0503" {
		t.Errorf("money (%d).Format(4) ==> string(%s)", b, b.Format(4))
	}
	if b.Format(10) != "23424234234.0503" {
		t.Errorf("money (%d).Format(10) ==> string(%s)", b, b.Format(4))
	}

	b = atype.NewMoney(-234242342340503)

	if b.Fmt(0) != "-23424234234" {
		t.Errorf("money (%d).FmtPercent(0) ==> string(%s)", b, b.Fmt(0))
	}
	if b.Fmt(1) != "-23424234234" {
		t.Errorf("money (%d).FmtPercent(1) ==> string(%s)", b, b.Fmt(1))
	}
	if b.Fmt(2) != "-23424234234.05" {
		t.Errorf("money (%d).FmtPercent(2) ==> string(%s)", b, b.Fmt(2))
	}
	if b.Fmt(3) != "-23424234234.05" {
		t.Errorf("money (%d).FmtPercent(3) ==> string(%s)", b, b.Fmt(3))
	}
	if b.Fmt(4) != "-23424234234.0503" {
		t.Errorf("money (%d).FmtPercent(4) ==> string(%s)", b, b.Fmt(4))
	}
	if b.Fmt(10) != "-23424234234.0503" {
		t.Errorf("money (%d).FmtPercent(10) ==> string(%s)", b, b.Fmt(4))
	}

}
func TestMoneyCalc(t *testing.T) {
	a := atype.ToMoney(500)
	b := atype.ToMoney(300)
	if (a.Add(b)) != atype.ToMoney(800) {
		t.Errorf("money 500+300 !=800")
	}
	if (a.Sub(b)) != atype.ToMoney(200) {
		t.Errorf("money 500-300 !=200")
	}
	if (a.Add(-b)) != atype.ToMoney(200) {
		t.Errorf("money 500-300 !=200")
	}
	if (a.Sub(-b)) != atype.ToMoney(800) {
		t.Errorf("money 500-300 !=200")
	}
}
func TestMoneyPanic(t *testing.T) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			t.Errorf("unknown panic %s", err)
		default:
			// well
		}
	}()
	m := atype.ToMoney(188.8)
	m.Sub(atype.MinMoney) // overflow

}
