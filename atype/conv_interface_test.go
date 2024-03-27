package atype_test

import (
	"encoding/json"
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

func TestConvStrings(t *testing.T) {
	s := `{"k":["a","b"]}`
	var x map[string]any
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		t.Error(err.Error())
	}
	b, ok := x["k"].([]any)
	if !ok {
		t.Error("parse strings fail")
	}

	ss := atype.ConvStrings(b)
	if len(ss) == 0 {
		t.Error("parse strings fail")
	}
	ss = atype.ConvStringsRaw(x["k"])
	if len(ss) == 0 {
		t.Error("parse strings fail")
	}

}
func TestConvStringMap(t *testing.T) {
	s := `{"k":{"a":"100","b":"200"}}`
	var x map[string]any
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		t.Error(err.Error())
	}
	_, ok := x["k"].(map[string]any)
	if !ok {
		t.Error("parse string map fail")
	}
}
func TestConvComplexStringMap(t *testing.T) {
	s := `{"k":{"a":{"b":"100"}}}`
	var x map[string]any
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		t.Error(err.Error())
	}
	b, ok := x["k"].(map[string]any)
	if !ok {
		t.Error("parse complex string map fail")
	}
	a := atype.ConvComplexStringMap(b)
	if len(a) == 0 {
		t.Error("parse complex string map fail")
	}
}
func TestConvStringsMap(t *testing.T) {
	s := `{"k":{"a":["100","200"]}}`
	var x map[string]any
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		t.Error(err.Error())
	}
	b, ok := x["k"].(map[string]any)
	if !ok {
		t.Error("parse complex string map fail")
	}
	a := atype.ConvStringsMap(b)
	if len(a) == 0 {
		t.Error("parse complex string map fail")
	}
}
func TestConvComplexStringsMap(t *testing.T) {
	s := `{"k":{"a":[["100","200"],["300","400","500"]]}}`
	var x map[string]any
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		t.Error(err.Error())
	}
	b, ok := x["k"].(map[string]any)
	if !ok {
		t.Error("parse complex string map fail")
	}
	a := atype.ConvComplexStringsMap(b)
	if len(a) == 0 {
		t.Error("parse complex string map fail")
	}
}
func TestConvStringMaps(t *testing.T) {
	s := `{"k":[{"a":"100"},{"b":"300","C":"400"}]}`
	var x map[string]any
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		t.Error(err.Error())
	}
	b, ok := x["k"].([]any)
	if !ok {
		t.Error("parse complex string map fail")
	}
	a := atype.ConvStringMaps(b)
	if len(a) == 0 {
		t.Error("parse complex string map fail")
	}
}
func TestConvComplexMaps(t *testing.T) {
	ss := []string{
		`{"a":100,"b":[{"得到": [{"士大夫": "算法撒旦"}, {"士大夫士大夫": "撒旦发射点"}]}]}`,
		`{"a":100,"b":[{"_TIP":"注释","得到": [{"士大夫": "算法撒旦"}, {"士大夫士大夫": "撒旦发射点"}]}]}`,
		`{"a":100,"b":[{"_TIP":"","得到": [{"士大夫": "算法撒旦", "_TIP":"注释"}, {"士大夫士大夫": "撒旦发射点"}]}]}`,
	}
	for _, s := range ss {
		var x map[string]any
		if err := json.Unmarshal([]byte(s), &x); err != nil {
			t.Error(err.Error())
		}

		b, ok := x["b"].([]any)
		if !ok {
			t.Error("parse json `b` fail " + s)
		}
		maps := atype.ConvComplexMaps(b)
		if len(maps) == 0 {
			t.Error("parse json `b` fail " + s)
		}
	}
}
