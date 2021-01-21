package util_test

import (
	"testing"

	"github.com/hi-iwi/AaGo/util"
)

func TestPrice(t *testing.T) {
	if f := util.NewPrice(2250).XPercent(0.14).Val(); f != 3.15 {
		t.Errorf("util.NewPrice(2250).XPercent(0.14) == %v, but it should be 3.15", f)
	}

	if f := util.NewPrice(2250).Mul(0.14).Val(); f != 315 {
		t.Errorf("util.NewPrice(2250).Mul(0.14) == %v, but it should be 315", f)
	}

	if f := util.NewPrice(315).Div(2250).Val(); f != 0.14 {
		t.Errorf("util.NewPrice(315).Mul(2250) == %v, but it should be 0.14", f)
	}

}
