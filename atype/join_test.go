package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

type stru struct {
	Name string `name:"name"`
	Age  int    `name:"age"`
}

func TestJoinTypeEnum(t *testing.T) {
	if atype.JoinValues&atype.JoinSortedBit != 0 {
		t.Errorf("atype.JoinValues&atype.JoinSortedBit != 0")
	}
	if atype.JoinKV&atype.JoinSortedBit != 0 {
		t.Errorf("atype.JoinKV&atype.JoinSortedBit != 0")
	}
	if atype.JoinJSON&atype.JoinSortedBit != 0 {
		t.Errorf("atype.JoinJSON&atype.JoinSortedBit != 0")
	}
	if atype.JoinMySQL&atype.JoinSortedBit != 0 {
		t.Errorf("atype.JoinMySQL&atype.JoinSortedBit != 0")
	}
	if atype.JoinURL&atype.JoinSortedBit != 0 {
		t.Errorf("atype.JoinURL&atype.JoinSortedBit != 0")
	}
	if atype.JoinSortedValues&atype.JoinSortedBit <= 0 {
		t.Errorf("atype.JoinSortedValues&atype.JoinSortedBit <= 0")
	}
	if atype.JoinSortedKV&atype.JoinSortedBit <= 0 {
		t.Errorf("atype.JoinUnsortedKV&atype.JoinSortedBit <= 0")
	}
	if atype.JoinSortedJSON&atype.JoinSortedBit <= 0 {
		t.Errorf("atype.JoinSortedJSON&atype.JoinSortedBit <= 0")
	}
	if atype.JoinSortedMySQL&atype.JoinSortedBit <= 0 {
		t.Errorf("atype.JoinSortedMySQL&atype.JoinSortedBit <= 0")
	}
	if atype.JoinSortedURL&atype.JoinSortedBit <= 0 {
		t.Errorf("atype.JoinSortedURL&atype.JoinSortedBit <= 0")
	}
}

func TestJoinByAlias(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := atype.JoinByNames(u, atype.JoinMySQL, " AND ", "name", "age")
	s1 := "`age`=\"18\" AND `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" AND `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("atype.JoinByAlias() == %s", s)
	}

}

func TestJoinAliasByElements(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := atype.JoinNamesByElements(u, atype.JoinMySQL, " AND ", "Name", "Age")
	s1 := "`age`=\"18\" AND `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" AND `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("atype.JoinNamesByElements() == %s", s)
	}

}
