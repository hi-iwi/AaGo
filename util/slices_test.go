package util_test

import (
	"fmt"
	"github.com/hi-iwi/AaGo/util"
	"golang.org/x/exp/slices"
	"testing"
)

func TestCompact(t *testing.T) {
	arr := []int{5, 1, 2, 2, 3, 4, 4, 5, 6, 7, 7, 8, 6}
	newArr := util.Compact(arr, false)
	// 此时 arr = [5 1 2 3 4 6 7 8 0 0 0 0 0]  newArr = [5 1 2 3 4 6 7 8]
	if !slices.Equal(newArr, []int{5, 1, 2, 3, 4, 6, 7, 8}) {
		t.Errorf("util.Compact() not passed")
	}

	arr = []int{5, 1, 2, 2, 3, 4, 4, 5, 6, 7, 7, 8, 6}
	slices.Sort(arr)
	newArr = util.Compact(arr, false)
	// 此时 arr = [5 1 2 3 4 6 7 8 0 0 0 0 0]  newArr = [5 1 2 3 4 6 7 8]
	if !slices.Equal(newArr, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("util.Compact() not passed")
	}

	fmt.Println(arr, newArr)
}
