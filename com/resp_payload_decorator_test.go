package com_test

import (
	"encoding/json"
	"fmt"
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"github.com/hi-iwi/AaGo/com"
	"testing"
)

func jsons(v any, e *ae.Error) string {
	if e != nil {
		fmt.Println(e.Text())
		return ""
	}
	s, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(s)
}

func TestStringifyPayloadFields(t *testing.T) {
	type Child struct {
		Test string `json:"test"`
	}
	type y struct {
		Child
		Bad    atype.Images   `json:"bad"`
		Tmp    atype.ImgSrc   `json:"-"`
		t      []atype.ImgSrc `json:"images"`
		Y      *int           `json:"y"`
		Img    *atype.ImgSrc  `json:"img"`
		Images []atype.ImgSrc `json:"ims"`
	}
	type x struct {
		Child
		A string       `json:"a"`
		B int64        `json:"b"`
		C int          `json:"c"`
		Y *y           `json:"y"`
		M *atype.Money `json:"m"`
		N atype.Money  `json:"money"`
	}
	img := atype.ImgSrc{}
	ims := []atype.ImgSrc{img}
	yy := 10000
	y0 := y{Y: &yy, Images: ims}
	a := x{A: "LOVE", B: 100, C: 300, Y: &y0}
	as := `{"a":"LOVE","b":"100","c":300,"m":null,"money":"0","test":"","y":{"bad":{},"img":null,"ims":[{"allowed":null,"crop_pattern":"","filetype":0,"height":0,"origin":"","path":"","provider":0,"resize_pattern":"","size":0,"width":0}],"test":"","y":10000}}`
	s := jsons(com.StringifyPayloadFields(a, "json"))
	if s != as {
		t.Errorf("%s --> %s", s, as)
	}
	b := []x{a}
	bs := "[" + as + "]"
	s = jsons(com.StringifyPayloadFields(b, "json"))
	if s != bs {
		t.Errorf("%s --> %s", s, bs)
	}
}
