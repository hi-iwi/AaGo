package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"runtime"
	"testing"
)

func TestPercent(t *testing.T) {
	p := atype.ToPercent(0.7)
	if p.Value().String() != "0.007" {
		t.Errorf("atype.NewPercent(0.7%%) : %s != 0.007", p.Value().String())
	}
	if p.Percent().String() != "0.7" {
		t.Errorf("atype.NewPercent(0.7%%) : %s != 0.7", p.Percent().String())
	}
}

func TestAmount(t *testing.T) {
	m := atype.ToAmount(188.8)
	if m.Fmt() != "188.8" {
		t.Errorf("atype.Umoney(1888000).Fmt() : %s != 188.8", m.Fmt())
	}
	if m.Format() != "188.80" {
		t.Errorf("atype.Umoney(1888000).Fmt() : %s != 188.80", m.Format())
	}
	p := atype.ToPercent(0.7)
	s := atype.ToAmount(2360).MulPercent(p).Fmt()
	if s != "16.52" {
		t.Errorf("2360*0.7%% : %s != 16.52", s)
	}

	b := atype.NewAmount(234242342340503)

	if b.Format(0) != "23424234234" {
		t.Errorf("amount (%d).Format(0) ==> string(%s)", b, b.Format(0))
	}
	if b.Format(1) != "23424234234.0" {
		t.Errorf("amount (%d).Format(1) ==> string(%s)", b, b.Format(1))
	}
	if b.Format(2) != "23424234234.05" {
		t.Errorf("amount (%d).Format(2) ==> string(%s)", b, b.Format(2))
	}
	if b.Format(3) != "23424234234.050" {
		t.Errorf("amount (%d).Format(3) ==> string(%s)", b, b.Format(3))
	}
	if b.Format(4) != "23424234234.0503" {
		t.Errorf("amount (%d).Format(4) ==> string(%s)", b, b.Format(4))
	}
	if b.Format(10) != "23424234234.0503" {
		t.Errorf("amount (%d).Format(10) ==> string(%s)", b, b.Format(4))
	}

	b = atype.NewAmount(-234242342340503)

	if b.Fmt(0) != "-23424234234" {
		t.Errorf("amount (%d).Fmt(0) ==> string(%s)", b, b.Fmt(0))
	}
	if b.Fmt(1) != "-23424234234" {
		t.Errorf("amount (%d).Fmt(1) ==> string(%s)", b, b.Fmt(1))
	}
	if b.Fmt(2) != "-23424234234.05" {
		t.Errorf("amount (%d).Fmt(2) ==> string(%s)", b, b.Fmt(2))
	}
	if b.Fmt(3) != "-23424234234.05" {
		t.Errorf("amount (%d).Fmt(3) ==> string(%s)", b, b.Fmt(3))
	}
	if b.Fmt(4) != "-23424234234.0503" {
		t.Errorf("amount (%d).Fmt(4) ==> string(%s)", b, b.Fmt(4))
	}
	if b.Fmt(10) != "-23424234234.0503" {
		t.Errorf("amount (%d).Fmt(10) ==> string(%s)", b, b.Fmt(4))
	}

}
func TestAmountCalc(t *testing.T) {
	a := atype.ToAmount(500)
	b := atype.ToAmount(300)
	if (a.Add(b)) != atype.ToAmount(800) {
		t.Errorf("amount 500+300 !=800")
	}
	if (a.Minus(b)) != atype.ToAmount(200) {
		t.Errorf("amount 500-300 !=200")
	}
	if (a.Add(-b)) != atype.ToAmount(200) {
		t.Errorf("amount 500-300 !=200")
	}
	if (a.Minus(-b)) != atype.ToAmount(800) {
		t.Errorf("amount 500-300 !=200")
	}
}
func TestAmountPanic(t *testing.T) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			t.Errorf("unknown panic %s", err)
		default:
			// well
		}
	}()
	m := atype.ToAmount(188.8)
	m.Minus(atype.MinAmount) // overflow

}
