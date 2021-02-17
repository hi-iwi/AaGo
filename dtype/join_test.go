package dtype_test

import (
	"testing"
)

type stru struct {
	Name string `alias:"name"`
	Age  int    `alias:"age"`
}

func TestJoinTypeEnum(t *testing.T) {
	if JoinValues&JoinSortedBit != 0 {
		t.Errorf("dtype.JoinValues&dtype.JoinSortedBit != 0")
	}
	if JoinKV&JoinSortedBit != 0 {
		t.Errorf("dtype.JoinKV&dtype.JoinSortedBit != 0")
	}
	if JoinJSON&JoinSortedBit != 0 {
		t.Errorf("dtype.JoinJSON&dtype.JoinSortedBit != 0")
	}
	if JoinMySQL&JoinSortedBit != 0 {
		t.Errorf("dtype.JoinMySQL&dtype.JoinSortedBit != 0")
	}
	if JoinURL&JoinSortedBit != 0 {
		t.Errorf("dtype.JoinURL&dtype.JoinSortedBit != 0")
	}
	if JoinSortedValues&JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedValues&dtype.JoinSortedBit <= 0")
	}
	if JoinSortedKV&JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinUnsortedKV&dtype.JoinSortedBit <= 0")
	}
	if JoinSortedJSON&JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedJSON&dtype.JoinSortedBit <= 0")
	}
	if JoinSortedMySQL&JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedMySQL&dtype.JoinSortedBit <= 0")
	}
	if JoinSortedURL&JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedURL&dtype.JoinSortedBit <= 0")
	}
}

func TestJoinByAlias(t *testing.T) {
	u := stru{
		Name: "Aario",
		Age:  18,
	}
	s := JoinByAlias(u, JoinMySQL, " AND ", "name", "age")
	if s != "`age`=\"18\" AND `name`=\"Aario\"" {
		t.Errorf("dtype.JoinByAlias() == %s", s)
	}

}

func TestJoinAliasByElements(t *testing.T) {
	u := stru{
		Name: "Aario",
		Age:  18,
	}
	s := JoinAliasByElements(u, JoinMySQL, " AND ", "Name", "Age")
	if s != "`age`=\"18\" AND `name`=\"Aario\"" {
		t.Errorf("dtype.JoinAliasByElements() == %s", s)
	}

}
