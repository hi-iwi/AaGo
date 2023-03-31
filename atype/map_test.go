package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

func TestAtypeMap(t *testing.T) {
	arr := map[interface{}]interface{}{
		1:      100,
		"name": "Iwi",
		"1":    "999",
		"test": map[string]interface{}{
			"nation": "China",
			"city":   "Shenzhen",
		},
		2: map[string]string{
			"sex": "male",
		},
	}
	x, _ := atype.NewMap(arr).Get("name")
	name := atype.String(x)

	t.Log("[\"name\"]", name)

	if name != "Iwi" {
		t.Error("[\"name\"] != Iwi")
	}
	y, _ := atype.NewMap(arr).Get(1)
	v, err := atype.Int(y)

	if v != 100 {
		t.Error("[1] != 100")
	} else {
		t.Logf(`[1] == %d %s`, v, err)
	}
	z, _ := atype.NewMap(arr).Get("1")
	if atype.String(z) != "999" {
		t.Error("[\"1\"] != 999")
	} else {
		w, _ := atype.NewMap(arr).Get("1")
		t.Logf("[\"1\"] == %s", atype.String(w))
	}
	h, _ := atype.NewMap(arr).Get("test.nation")
	nation := atype.String(h)
	if nation != "China" {
		t.Log("[\"test\".\"nation\"] != China")
	} else {
		t.Log("[\"test\".\"nation\"]", nation)
	}

	t.Log(atype.NewMap(arr).Get(2, "sex"))
	o, _ := atype.NewMap(arr).Get(2, "sex")
	sex := atype.String(o)
	t.Logf("[2.\"sex\"] == %s", sex)
}
