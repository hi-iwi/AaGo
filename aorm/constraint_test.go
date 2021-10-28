package aorm_test

import (
	"github.com/hi-iwi/AaGo/aorm"
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
	s := aorm.And(u, "name", "age")
	s1 := "`age`=\"18\" AND `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" AND `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("aorm.And(u, true, ...) == %s", s)
	}

	s = aorm.And(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("aorm.And(u, false, ...) == %s", s)
	}
}

func TestOr(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := aorm.Or(u, "name", "age")
	s1 := "`age`=\"18\" OR `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" OR `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("aorm.Or(u, true, ...) == %s", s)
	}

	s = aorm.Or(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("aorm.Or(u, false, ...) == %s", s)
	}
}

func TestAndWithWhere(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := aorm.AndWithWhere(u, "name", "age")
	s1 := " WHERE `age`=\"18\" AND `name`=\"Iwi\" "
	s2 := " WHERE `name`=\"Iwi\" AND `age`=\"18\" "
	if s != s1 && s != s2 {
		t.Errorf("aorm.Or(u, true, ...) == %s", s)
	}

	s = aorm.AndWithWhere(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("aorm.Or(u, false, ...) == %s", s)
	}
}

func TestOrWithWhere(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := aorm.OrWithWhere(u, "name", "age")
	s1 := " WHERE `age`=\"18\" OR `name`=\"Iwi\" "
	s2 := " WHERE `name`=\"Iwi\" OR `age`=\"18\" "
	if s != s1 && s != s2 {
		t.Errorf("aorm.Or(u, true, ...) == %s", s)
	}

	s = aorm.OrWithWhere(u, "Name", "Age")
	if s != s1 && s != s2 {
		t.Errorf("aorm.Or(u, false, ...) == %s", s)
	}
}
