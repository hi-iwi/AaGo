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

	name := atype.String(atype.NewMap(arr).Get("name"))

	t.Log("[\"name\"]", name)

	if name != "Iwi" {
		t.Error("[\"name\"] != Iwi")
	}

	v, err := atype.Int(atype.NewMap(arr).Get(1))

	if v != 100 {
		t.Error("[1] != 100")
	} else {
		t.Logf(`[1] == %d %s`, v, err)
	}

	if atype.String(atype.NewMap(arr).Get("1")) != "999" {
		t.Error("[\"1\"] != 999")
	} else {
		t.Logf("[\"1\"] == %s", atype.String(atype.NewMap(arr).Get("1")))
	}

	nation := atype.String(atype.NewMap(arr).Get("test.nation"))
	if nation != "China" {
		t.Log("[\"test\".\"nation\"] != China")
	} else {
		t.Log("[\"test\".\"nation\"]", nation)
	}

	t.Log(atype.NewMap(arr).Get(2, "sex"))
	sex := atype.String(atype.NewMap(arr).Get(2, "sex"))
	t.Logf("[2.\"sex\"] == %s", sex)
}
