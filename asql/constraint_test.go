package asql_test

import (
	"github.com/hi-iwi/AaGo/asql"
	"testing"
)

type stru struct {
	Name string `name:"name"`
	Age  int    `name:"age"`
}

func TestAnd(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := asql.And(u, "name", "age")
	s1 := "`age`=\"18\" AND `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" AND `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("asql.And(u, true, ...) == %s", s)
	}

	s = asql.And(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("asql.And(u, false, ...) == %s", s)
	}
}

func TestOr(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := asql.Or(u, "name", "age")
	s1 := "`age`=\"18\" OR `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" OR `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("asql.Or(u, true, ...) == %s", s)
	}

	s = asql.Or(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("asql.Or(u, false, ...) == %s", s)
	}
}

func TestAndWithWhere(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := asql.AndWithWhere(u, "name", "age")
	s1 := " WHERE `age`=\"18\" AND `name`=\"Iwi\" "
	s2 := " WHERE `name`=\"Iwi\" AND `age`=\"18\" "
	if s != s1 && s != s2 {
		t.Errorf("asql.Or(u, true, ...) == %s", s)
	}

	s = asql.AndWithWhere(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("asql.Or(u, false, ...) == %s", s)
	}
}

func TestOrWithWhere(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := asql.OrWithWhere(u, "name", "age")
	s1 := " WHERE `age`=\"18\" OR `name`=\"Iwi\" "
	s2 := " WHERE `name`=\"Iwi\" OR `age`=\"18\" "
	if s != s1 && s != s2 {
		t.Errorf("asql.Or(u, true, ...) == %s", s)
	}

	s = asql.OrWithWhere(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("asql.Or(u, false, ...) == %s", s)
	}
}
