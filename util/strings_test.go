package util_test

import (
	"github.com/hi-iwi/AaGo/util"
	"testing"
)

func TestIndexRunes(t *testing.T) {
	content := "你好，我是一苇Iwi。"
	if p := util.IndexRunes([]rune(content), []rune("你")); p != 0 {
		t.Errorf("util.IndexRunes error indexof `你` %d", p)
	}
	if p := util.IndexRunes([]rune(content), []rune("一苇")); p != 5 {
		t.Errorf("util.IndexRunes error indexof `一苇` %d", p)
	}
	if p := util.IndexRunes([]rune(content), []rune("w")); p != 8 {
		t.Errorf("util.IndexRunes error indexof `w` %d", p)
	}
	if p := util.IndexRunes([]rune(content), []rune("i。")); p != 9 {
		t.Errorf("util.IndexRunes error indexof `i。` %d", p)
	}
	if p := util.IndexRunes([]rune(content), []rune("。")); p != 10 {
		t.Errorf("util.IndexRunes error indexof `。` %d", p)
	}
}
