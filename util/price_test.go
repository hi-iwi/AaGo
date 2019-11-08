package util_test

import (
	"testing"

	"github.com/luexu/AaGo/util"
)

func TestPrice(t *testing.T) {
	p := util.NewPrice(2250)
	if f := p.XPercent(0.14).Val(); f != 3.15 {
		t.Errorf("util.NewPrice(2250).XPercent(0.14) == %v, but it should be 3.15", f)
	}
}
