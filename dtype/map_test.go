package dtype_test

import (
	"testing"
)

func TestDtypeMap(t *testing.T) {
	arr := map[interface{}]interface{}{
		1:      100,
		"name": "Aario",
		"1":    "999",
		"test": map[string]interface{}{
			"nation": "China",
			"city":   "Shenzhen",
		},
		2: map[string]string{
			"sex": "male",
		},
	}

	name := String(NewMap(arr).Get("name"))

	t.Log("[\"name\"]", name)

	if name != "Aario" {
		t.Error("[\"name\"] != Aario")
	}

	v, err := Int(NewMap(arr).Get(1))

	if v != 100 {
		t.Error("[1] != 100")
	} else {
		t.Logf(`[1] == %d %s`, v, err)
	}

	if String(NewMap(arr).Get("1")) != "999" {
		t.Error("[\"1\"] != 999")
	} else {
		t.Logf("[\"1\"] == %s", String(NewMap(arr).Get("1")))
	}

	nation := String(NewMap(arr).Get("test.nation"))
	if nation != "China" {
		t.Log("[\"test\".\"nation\"] != China")
	} else {
		t.Log("[\"test\".\"nation\"]", nation)
	}

	t.Log(NewMap(arr).Get(2, "sex"))
	sex := String(NewMap(arr).Get(2, "sex"))
	t.Logf("[2.\"sex\"] == %s", sex)
}
