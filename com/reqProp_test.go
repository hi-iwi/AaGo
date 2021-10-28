package com_test

import (
	"testing"

	"github.com/hi-iwi/AaGo/com"
)

func TestSetRegExp(t *testing.T) {
	type vt uint
	var a, b vt
	a = 100
	b = 200
	exp := com.UintsRegExp(a, b)
	if exp != "^(100|200)$" {
		t.Error("Uint64RegExp error: " + exp)
	}
}
func TestReqPropFilter(t *testing.T) {
	p := com.NewReqProp("test", "Iwi")
	if p.IsEmpty() {
		t.Error("test='Iwi' should be not empty")
	}
	t.Log("string('Iwi') ==> ", p.String())
	if i, b := p.Int(); b == nil {
		t.Error("'Iwi' should not can be converted into int")
	} else {
		t.Log("int('Iwi') ==> ", i)
	}
	if f, b := p.Float64(); b == nil {
		t.Error("'Iwi' should not can be converted into float64")
	} else {
		t.Log("float64('Iwi') ==> ", f)
	}
	if err := p.Filter(`\d+`, true); err == nil {
		t.Error("'Iwi' should not match `\\d+`")
	} else {
		t.Log("'Iwi' doesn't match `\\d+`")
	}
	if err := p.Filter(`[[:word:]]+`, true); err != nil {
		t.Error("'Iwi' should match `[[:word:]]`")
	} else {
		t.Log("'Iwi' matches `[[:word:]]`")
	}

	t.Log("----------------")

	p = com.NewReqProp("test", "100")
	if !p.NotEmpty() {
		t.Error("test='100' should be not empty")
	}
	t.Log("string('100') ==> ", p.String())
	if i, b := p.Int(); b != nil {
		t.Error("'100' should can be converted into int")
	} else {
		t.Log("int('100') ==> ", i)
	}
	if f, b := p.Float64(); b != nil {
		t.Error("'100' should can be converted into float64")
	} else {
		t.Log("float64('100') ==> ", f)
	}
	if err := p.Filter(`\d+`, true); err != nil {
		t.Error("'100' should match `\\d+`", err)
	} else {
		t.Log("'100' doesn't match `\\d+`")
	}
	if err := p.Filter(`[[:word:]]+`, true); err != nil {
		t.Error("'100' should match `[[:word:]]`", err)
	} else {
		t.Log("'100' matches `[[:word:]]`")
	}

	t.Log("----------------")

	p = com.NewReqProp("test", 9527)
	if !p.NotEmpty() {
		t.Error("test=9527 should be not empty")
	}
	t.Log("string(9527) ==> ", p.String())
	if i, b := p.Int(); b != nil {
		t.Error("9527 should can be converted into int")
	} else {
		t.Log("int(9527) ==> ", i)
	}
	if f, b := p.Float64(); b != nil {
		t.Error("9527 should can be converted into float64")
	} else {
		t.Log("float64(9527) ==> ", f)
	}
	if err := p.Filter(`\d+`, true); err != nil {
		t.Error("9527 should match `\\d+`", err)
	} else {
		t.Log("9527 matches `\\d+`")
	}
	if err := p.Filter(`a\d+`, true); err == nil {
		t.Error("9527 should not match `a\\d+`")
	} else {
		t.Log("9527 doesn't matches `a\\d+`", err)
	}

	t.Log("----------------")

	p = com.NewReqProp("test", "")
	if p.NotEmpty() {
		t.Error("test='' should be empty")
	} else {
		t.Log("test='' is empty")
	}
	t.Log("string('') ==> ", p.String())
	if i, b := p.Int(); b == nil {
		t.Error("'' should can't be converted into int")
	} else {
		t.Log("int('') ==> ", i, b)
	}
	if f, b := p.Float64(); b == nil {
		t.Error("'' should can't be converted into float64")
	} else {
		t.Log("float64('') ==> ", f, b)
	}
	if err := p.Filter(true); err == nil {
		t.Error("'' should be filtered when it's required", err)
	} else {
		t.Log("'' is filtered because it's required")
	}

	if err := p.Filter(`\w+`); err == nil {
		t.Error("'' should be filtered by `\\w+`", err)
	}
}

// func BenchmarkReqPropFilter(t *testing.B) {
// 	t.StopTimer()  // 停止计时
// 	t.StartTimer() // 开始计时
// 	for i := 0; i < t.N; i++ {

// 	}
// }
