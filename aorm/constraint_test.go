package aorm_test

import (
	"testing"
)

type stru struct {
	Name string `alias:"name"`
	Age  int    `alias:"age"`
}

func TestAnd(t *testing.T) {
	u := stru{
		Name: "Aario",
		Age:  18,
	}
	s := And(u, "name", "age")
	if s != "`age`=\"18\" AND `name`=\"Aario\"" {
		t.Errorf("aorm.And(u, true, ...) == %s", s)
	}

	s = And(u, "Name", "Age")
	if s != "`age`=\"18\" AND `name`=\"Aario\"" {
		t.Errorf("aorm.And(u, false, ...) == %s", s)
	}
}

func TestOr(t *testing.T) {
	u := stru{
		Name: "Aario",
		Age:  18,
	}
	s := Or(u, "name", "age")
	if s != "`age`=\"18\" OR `name`=\"Aario\"" {
		t.Errorf("aorm.Or(u, true, ...) == %s", s)
	}

	s = Or(u, "Name", "Age")
	if s != "`age`=\"18\" OR `name`=\"Aario\"" {
		t.Errorf("aorm.Or(u, false, ...) == %s", s)
	}
}

func TestAndWithWhere(t *testing.T) {
	u := stru{
		Name: "Aario",
		Age:  18,
	}
	s := AndWithWhere(u, "name", "age")
	if s != " WHERE `age`=\"18\" AND `name`=\"Aario\" " {
		t.Errorf("aorm.Or(u, true, ...) == %s", s)
	}

	s = AndWithWhere(u, "Name", "Age")
	if s != " WHERE `age`=\"18\" AND `name`=\"Aario\" " {
		t.Errorf("aorm.Or(u, false, ...) == %s", s)
	}
}

func TestOrWithWhere(t *testing.T) {
	u := stru{
		Name: "Aario",
		Age:  18,
	}
	s := OrWithWhere(u, "name", "age")
	if s != " WHERE `age`=\"18\" OR `name`=\"Aario\" " {
		t.Errorf("aorm.Or(u, true, ...) == %s", s)
	}

	s = OrWithWhere(u, "Name", "Age")
	if s != " WHERE `age`=\"18\" OR `name`=\"Aario\" " {
		t.Errorf("aorm.Or(u, false, ...) == %s", s)
	}
}
