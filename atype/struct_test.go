package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

type congestion struct {
	Iwi string `name:"iwi" validation:"min=2,max=32"`
	Tom string `name:"tom"`
}

func TestTag(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
	}

	if name := atype.Tag(c, "Iwi", "name"); name != "iwi" {
		t.Errorf("congestion struct Iwi's name is iwi, not %s", name)
	}

	if vali := atype.Tag(c, "Iwi", "validation"); vali == "" {
		t.Errorf("congestion struct Iwi's validation is not empty")
	}
}
func TestAliasTag(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
	}

	if name := atype.NameTag(c, "Iwi"); name != "iwi" {
		t.Errorf("congestion struct Iwi's name is iwi, not %s", name)
	}

	if vali := atype.NameTag(c, "Iwi"); vali == "" {
		t.Errorf("congestion struct Iwi's validation is not empty")
	}
}

func TestValueByTag(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
		Tom: "This's Tom",
	}
	if v, err := atype.ValueByTag(c, "name", "iwi"); err != nil {
		t.Errorf("atype.ValueByTag error")
	} else {
		if x := atype.String(v); x != c.Iwi {
			t.Errorf("atype.ValueByTag is not matched. %s != %s", x, c.Iwi)
		}
	}

	if v, err := atype.ValueByTag(c, "name", "tom"); err != nil {
		t.Errorf("atype.ValueByTag error")
	} else {
		if x := atype.String(v); x != c.Tom {
			t.Errorf("atype.ValueByTag is not matched. %s != %s", x, c.Tom)
		}
	}
}

func TestValueByAlias(t *testing.T) {
	c := congestion{
		Iwi: "Hi, Iwi",
		Tom: "This's Tom",
	}
	if v, err := atype.ValueByName(c, "iwi"); err != nil {
		t.Errorf("atype.ValueByTag error")
	} else {
		if x := atype.String(v); x != c.Iwi {
			t.Errorf("atype.ValueByTag is not matched. %s != %s", x, c.Iwi)
		}
	}

	if v, err := atype.ValueByName(c, "tom"); err != nil {
		t.Errorf("atype.ValueByTag error")
	} else {
		if x := atype.String(v); x != c.Tom {
			t.Errorf("atype.ValueByTag is not matched. %s != %s", x, c.Tom)
		}
	}
}
