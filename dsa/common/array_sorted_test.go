package common

import (
	"testing"
)

var arrSortedTestData = []struct {
	input  []int
	expRes bool
}{
	{input: []int{1, 2, 3, 4, 5, 7}, expRes: true},
	{input: []int{7}, expRes: true},
	{input: []int{1, 2, 5, 4}, expRes: false},
	{input: []int{5, 4, 3, 2, 1}, expRes: false},
}

func TestIsArraySorted(t *testing.T) {
	for _, td := range arrSortedTestData {
		res := IsArraySorted(td.input)
		if td.expRes != res {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
