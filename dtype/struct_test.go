package dtype_test

import (
	"testing"
)

type congestion struct {
	Aario string `alias:"aario" validation:"min=2,max=32"`
	Tom   string `alias:"tom"`
}

func TestTag(t *testing.T) {
	c := congestion{
		Aario: "Hi, Aario",
	}

	if alias := Tag(c, "Aario", "alias"); alias != "aario" {
		t.Errorf("congestion struct Aario's alias is aario, not %s", alias)
	}

	if vali := Tag(c, "Aario", "validation"); vali == "" {
		t.Errorf("congestion struct Aario's validation is not empty")
	}
}
func TestAliasTag(t *testing.T) {
	c := congestion{
		Aario: "Hi, Aario",
	}

	if alias := AliasTag(c, "Aario"); alias != "aario" {
		t.Errorf("congestion struct Aario's alias is aario, not %s", alias)
	}

	if vali := AliasTag(c, "Aario"); vali == "" {
		t.Errorf("congestion struct Aario's validation is not empty")
	}
}

func TestValueByTag(t *testing.T) {
	c := congestion{
		Aario: "Hi, Aario",
		Tom:   "This's Tom",
	}
	if v, err := ValueByTag(c, "alias", "aario"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := String(v); x != c.Aario {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Aario)
		}
	}

	if v, err := ValueByTag(c, "alias", "tom"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := String(v); x != c.Tom {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Tom)
		}
	}
}

func TestValueByAlias(t *testing.T) {
	c := congestion{
		Aario: "Hi, Aario",
		Tom:   "This's Tom",
	}
	if v, err := ValueByAlias(c, "aario"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := String(v); x != c.Aario {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Aario)
		}
	}

	if v, err := ValueByAlias(c, "tom"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := String(v); x != c.Tom {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Tom)
		}
	}
}
