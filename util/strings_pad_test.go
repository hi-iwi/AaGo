package util_test

import (
	"github.com/hi-iwi/AaGo/util"
	"testing"
)

func TestPad(t *testing.T) {
	var s = "Aario"
	if util.Pad(s, " ", len(s)-1) != s {
		t.Errorf("util.Pad(%s, %s, %d) not passed", s, " ", len(s)-1)
	}
	if util.Unpad(s, " ") != s {
		t.Errorf("util.Unpad(%s, %s) not passed", s, " ")
	}

	if util.Pad(s, " ", len(s)+1) != " "+s {
		t.Errorf("util.Pad(%s, %s, %d) not passed", s, " ", len(s)+1)
	}
	if util.Unpad(" "+s, " ") != s {
		t.Errorf("util.Unpad(%s, %s) not passed", s, " ")
	}

	s2 := util.Pad(s, "^_^", 20)
	if len(s2) < 20 {
		t.Errorf("util.Pad(%s, %s, %d) not passed", s, "^_^", 20)
	}
	if util.Unpad(s2, "^_^") != s {
		t.Errorf("util.Unpad(%s, %s) not passed", s2, "^_^")
	}
}
func TestPadRight(t *testing.T) {
	var s = "Aario"
	if util.PadRight(s, " ", len(s)-1) != s {
		t.Errorf("util.Pad(%s, %s, %d) not passed", s, " ", len(s)-1)
	}
	if util.UnpadRight(s, " ") != s {
		t.Errorf("util.Unpad(%s, %s) not passed", s, " ")
	}

	if util.PadRight(s, " ", len(s)+1) != s+" " {
		t.Errorf("util.Pad(%s, %s, %d) not passed", s, " ", len(s)+1)
	}
	if util.UnpadRight(s+" ", " ") != s {
		t.Errorf("util.Unpad(%s, %s) not passed", s, " ")
	}

	s2 := util.PadRight(s, "^_^", 20)
	if len(s2) < 20 {
		t.Errorf("util.Pad(%s, %s, %d) not passed", s, "^_^", 20)
	}
	if util.UnpadRight(s2, "^_^") != s {
		t.Errorf("util.Unpad(%s, %s) not passed", s2, "^_^")
	}
}
