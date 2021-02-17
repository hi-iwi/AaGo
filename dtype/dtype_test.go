package dtype_test

import (
	"testing"
)

func TestDtype(t *testing.T) {
	b := 234242342342423.3
	s := New(b).String()
	if s != "234242342342423.3" {
		t.Errorf("float64(%f) ==> string(%s)", b, s)
	}

}
func TestDtypeGet(t *testing.T) {
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

	d := New(arr)
	v, err := d.Get("name")
	t.Log("[\"name\"]", v, err)

	if v.String() != "Aario" {
		t.Error("[\"name\"] != Aario")
	}

	v, err = d.Get(1)
	t.Log("[1]", v, err)

	i, err := v.Int()
	if i != 100 {
		t.Error("[1] != 100")
	}

	v, err = d.Get("1")
	t.Log("[\"1\"]", v, err)
	if v.String() != "999" {
		t.Error("[\"1\"] != 999")
	}

	v, err = d.Get("test.nation")
	t.Log("[\"test\".\"nation\"]", v, err)

	v, err = d.Get(2, "sex")
	t.Log("[2.\"sex\"]", v, err)
}
