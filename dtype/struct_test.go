package dtype_test

import (
	"github.com/hi-iwi/AaGo/dtype"
	"testing"
)

type congestion struct {
	Iwi string `name:"iwi" validation:"min=2,max=32"`
	Tom   string `name:"tom"`
}

func TestTag(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
	}

	if name := dtype.Tag(c, "Iwi", "name"); name != "iwi" {
		t.Errorf("congestion struct Iwi's name is iwi, not %s", name)
	}

	if vali := dtype.Tag(c, "Iwi", "validation"); vali == "" {
		t.Errorf("congestion struct Iwi's validation is not empty")
	}
}
func TestAliasTag(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
	}

	if name := dtype.NameTag(c, "Iwi"); name != "iwi" {
		t.Errorf("congestion struct Iwi's name is iwi, not %s", name)
	}

	if vali := dtype.NameTag(c, "Iwi"); vali == "" {
		t.Errorf("congestion struct Iwi's validation is not empty")
	}
}

func TestValueByTag(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
		Tom:   "This's Tom",
	}
	if v, err := dtype.ValueByTag(c, "name", "iwi"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := dtype.String(v); x != c.Iwi {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Iwi)
		}
	}

	if v, err := dtype.ValueByTag(c, "name", "tom"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := dtype.String(v); x != c.Tom {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Tom)
		}
	}
}

func TestValueByAlias(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
		Tom:   "This's Tom",
	}
	if v, err := dtype.ValueByName(c, "iwi"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := dtype.String(v); x != c.Iwi {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Iwi)
		}
	}

	if v, err := dtype.ValueByName(c, "tom"); err != nil {
		t.Errorf("dtype.ValueByTag error")
	} else {
		if x := dtype.String(v); x != c.Tom {
			t.Errorf("dtype.ValueByTag is not matched. %s != %s", x, c.Tom)
		}
	}
}
