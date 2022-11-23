package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

func TestAmountn(t *testing.T) {
	b := atype.Umoney(234242342340503)

	if b.Fmt(0) != "23424234234" {
		t.Errorf("amount (%d).Fmt(0) ==> string(%s)", b, b.Fmt(0))
	}
	if b.Fmt(1) != "23424234234.0" {
		t.Errorf("amount (%d).Fmt(1) ==> string(%s)", b, b.Fmt(1))
	}
	if b.Fmt(2) != "23424234234.05" {
		t.Errorf("amount (%d).Fmt(2) ==> string(%s)", b, b.Fmt(2))
	}
	if b.Fmt(3) != "23424234234.050" {
		t.Errorf("amount (%d).Fmt(3) ==> string(%s)", b, b.Fmt(3))
	}
	if b.Fmt(4) != "23424234234.0503" {
		t.Errorf("amount (%d).Fmt(4) ==> string(%s)", b, b.Fmt(4))
	}
	if b.Fmt(10) != "23424234234.0503" {
		t.Errorf("amount (%d).Fmt(10) ==> string(%s)", b, b.Fmt(4))
	}
}
func TestAmount(t *testing.T) {
	b := atype.Money(-234242342340503)

	if b.Fmt(0) != "-23424234234" {
		t.Errorf("amount (%d).Fmt(0) ==> string(%s)", b, b.Fmt(0))
	}
	if b.Fmt(1) != "-23424234234.0" {
		t.Errorf("amount (%d).Fmt(1) ==> string(%s)", b, b.Fmt(1))
	}
	if b.Fmt(2) != "-23424234234.05" {
		t.Errorf("amount (%d).Fmt(2) ==> string(%s)", b, b.Fmt(2))
	}
	if b.Fmt(3) != "-23424234234.050" {
		t.Errorf("amount (%d).Fmt(3) ==> string(%s)", b, b.Fmt(3))
	}
	if b.Fmt(4) != "-23424234234.0503" {
		t.Errorf("amount (%d).Fmt(4) ==> string(%s)", b, b.Fmt(4))
	}
	if b.Fmt(10) != "-23424234234.0503" {
		t.Errorf("amount (%d).Fmt(10) ==> string(%s)", b, b.Fmt(4))
	}
}
