package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"reflect"
	"testing"
)

func f0(i interface{}) reflect.Kind {
	return atype.PrimitiveType(&i)
}

// 获取原始类型
func TestPrimitiveType(t *testing.T) {
	type safeInt atype.Int24
	type b safeInt
	type c b
	x := c(100)
	t.Log(atype.PrimitiveType(&x), atype.PType(&x), atype.PType(x))
	type g1 struct {
		A int64 `json:"a"`
	}
	type g2 g1
	g := g2{A: 200}
	gg := &g
	gg2 := &gg

	t.Log(atype.PrimitiveType(&g), atype.PType(&g), atype.PType(g))
	t.Log(atype.PrimitiveType(&gg), atype.PType(&gg), atype.PType(gg))
	t.Log(atype.PrimitiveType(&gg2), f0(gg2), atype.PType(&gg2), atype.PType(gg2))

	type y struct {
		A string `json:"a"`
		B int64  `json:"b"`
		C int    `json:"c"`
	}
	a := y{A: "LOVE", B: 100, C: 300}
	t.Log(f0(&a), atype.PType(&a), atype.PType(a))
}
func TestPrimitiveType2(t *testing.T) {
	type y struct {
		Tmp atype.ImgSrc  `json:"-"`
		t   atype.Images  `json:"images"`
		Y   *int          `json:"y"`
		Img *atype.ImgSrc `json:"img"`
	}
	type x struct {
		A string `json:"a"`
		B int64  `json:"b"`
		C int    `json:"c"`
		Y *y     `json:"y"`
	}
	yy := 10000
	y0 := y{Y: &yy}
	a := x{A: "LOVE", B: 100, C: 300, Y: &y0}
	t.Log(atype.PType(a.Y.Img))
}
