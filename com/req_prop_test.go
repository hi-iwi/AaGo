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
	if _, b := p.Int(); b == nil {
		t.Error("'Iwi' should not can be converted into int")
	}
	if _, b := p.Float64(); b == nil {
		t.Error("'Iwi' should not can be converted into float64")
	}
	if err := p.Filter(`\d+`, true); err == nil {
		t.Error("'Iwi' should not match `\\d+`")
	}
	if err := p.Filter(`[[:word:]]+`, true); err != nil {
		t.Error("'Iwi' should match `[[:word:]]`")
	}

	p = com.NewReqProp("test", "100")
	if !p.NotEmpty() {
		t.Error("test='100' should be not empty")
	}
	if _, b := p.Int(); b != nil {
		t.Error("'100' should can be converted into int")
	}
	if _, b := p.Float64(); b != nil {
		t.Error("'100' should can be converted into float64")
	}
	if err := p.Filter(`\d+`, true); err != nil {
		t.Error("'100' should match `\\d+`", err)
	}
	if err := p.Filter(`[[:word:]]+`, true); err != nil {
		t.Error("'100' should match `[[:word:]]`", err)
	}

	p = com.NewReqProp("test", 9527)
	if !p.NotEmpty() {
		t.Error("test=9527 should be not empty")
	}
	if _, b := p.Int(); b != nil {
		t.Error("9527 should can be converted into int")
	}
	if _, b := p.Float64(); b != nil {
		t.Error("9527 should can be converted into float64")
	}
	if err := p.Filter(`\d+`, true); err != nil {
		t.Error("9527 should match `\\d+`", err)
	}
	if err := p.Filter(`a\d+`, true); err == nil {
		t.Error("9527 should not match `a\\d+`")
	}

	p = com.NewReqProp("test", "")
	if p.NotEmpty() {
		t.Error("test='' should be empty")
	}
	if _, b := p.Int(); b == nil {
		t.Error("'' should can't be converted into int")
	}
	if _, b := p.Float64(); b == nil {
		t.Error("'' should can't be converted into float64")
	}
	if err := p.Filter(true); err == nil {
		t.Error("'' should be filtered when it's required", err)
	}

	if err := p.Filter(`\w+`); err == nil {
		t.Error("'' should be filtered by `\\w+`", err)
	}
}
 