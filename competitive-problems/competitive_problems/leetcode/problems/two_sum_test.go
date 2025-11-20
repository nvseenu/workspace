package problems

import (
	"slices"
	"testing"
)

var twoSumTestData = []struct {
	nums   []int
	target int
	expRes []int
}{
	{nums: []int{3, 2, 4}, target: 6, expRes: []int{1, 2}},
	{nums: []int{2, 7, 11, 15}, target: 9, expRes: []int{0, 1}},
}

func TestTwoSum(t *testing.T) {

	for _, td := range twoSumTestData {
		res := TwoSum(td.nums, td.target)
		if !slices.Equal(res, td.expRes) {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
