package dtype_test

import (
	"github.com/hi-iwi/AaGo/dtype"
	"testing"
)

type stru struct {
	Name string `name:"name"`
	Age  int    `name:"age"`
}

func TestJoinTypeEnum(t *testing.T) {
	if dtype.JoinValues&dtype.JoinSortedBit != 0 {
		t.Errorf("dtype.JoinValues&dtype.JoinSortedBit != 0")
	}
	if dtype.JoinKV&dtype.JoinSortedBit != 0 {
		t.Errorf("dtype.JoinKV&dtype.JoinSortedBit != 0")
	}
	if dtype.JoinJSON&dtype.JoinSortedBit != 0 {
		t.Errorf("dtype.JoinJSON&dtype.JoinSortedBit != 0")
	}
	if dtype.JoinMySQL&dtype.JoinSortedBit != 0 {
		t.Errorf("dtype.JoinMySQL&dtype.JoinSortedBit != 0")
	}
	if dtype.JoinURL&dtype.JoinSortedBit != 0 {
		t.Errorf("dtype.JoinURL&dtype.JoinSortedBit != 0")
	}
	if dtype.JoinSortedValues&dtype.JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedValues&dtype.JoinSortedBit <= 0")
	}
	if dtype.JoinSortedKV&dtype.JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinUnsortedKV&dtype.JoinSortedBit <= 0")
	}
	if dtype.JoinSortedJSON&dtype.JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedJSON&dtype.JoinSortedBit <= 0")
	}
	if dtype.JoinSortedMySQL&dtype.JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedMySQL&dtype.JoinSortedBit <= 0")
	}
	if dtype.JoinSortedURL&dtype.JoinSortedBit <= 0 {
		t.Errorf("dtype.JoinSortedURL&dtype.JoinSortedBit <= 0")
	}
}

func TestJoinByAlias(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := dtype.JoinByNames(u, dtype.JoinMySQL, " AND ", "name", "age")
	s1 := "`age`=\"18\" AND `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" AND `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("dtype.JoinByAlias() == %s", s)
	}

}

func TestJoinAliasByElements(t *testing.T) {
	u := stru{
		Name: "Iwi",
		Age:  18,
	}
	s := dtype.JoinNamesByElements(u, dtype.JoinMySQL, " AND ", "Name", "Age")
	s1 := "`age`=\"18\" AND `name`=\"Iwi\""
	s2 := "`name`=\"Iwi\" AND `age`=\"18\""
	if s != s1 && s != s2 {
		t.Errorf("dtype.JoinNamesByElements() == %s", s)
	}

}
