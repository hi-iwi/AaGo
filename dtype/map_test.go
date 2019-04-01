package dtype_test

import (
	"testing"

	"github.com/luexu/AaGo/dtype"
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

	name := dtype.String(dtype.NewMap(arr).Get("name"))

	t.Log("[\"name\"]", name)

	if name != "Aario" {
		t.Error("[\"name\"] != Aario")
	}

	v, err := dtype.Int(dtype.NewMap(arr).Get(1))

	if v != 100 {
		t.Error("[1] != 100")
	} else {
		t.Logf(`[1] == %d %s`, v, err)
	}

	if dtype.String(dtype.NewMap(arr).Get("1")) != "999" {
		t.Error("[\"1\"] != 999")
	} else {
		t.Logf("[\"1\"] == %s", dtype.String(dtype.NewMap(arr).Get("1")))
	}

	nation := dtype.String(dtype.NewMap(arr).Get("test.nation"))
	if nation != "China" {
		t.Log("[\"test\".\"nation\"] != China")
	} else {
		t.Log("[\"test\".\"nation\"]", nation)
	}

	t.Log(dtype.NewMap(arr).Get(2, "sex"))
	sex := dtype.String(dtype.NewMap(arr).Get(2, "sex"))
	t.Logf("[2.\"sex\"] == %s", sex)
}
