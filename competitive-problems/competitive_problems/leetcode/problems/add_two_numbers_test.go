package problems

import (
	"slices"
	"testing"
)

var addTwoNumTestData = []struct {
	ls1    []int
	ls2    []int
	expRes []int
}{
	{ls1: []int{2, 4, 3}, ls2: []int{5, 6, 4}, expRes: []int{7, 0, 8}},
	{ls1: []int{0}, ls2: []int{0}, expRes: []int{0}},
	{ls1: []int{9, 9, 9, 9, 9, 9, 9}, ls2: []int{9, 9, 9, 9}, expRes: []int{8, 9, 9, 9, 0, 0, 0, 1}},
}

func TestAddNums(t *testing.T) {

	for _, td := range addTwoNumTestData {
		ls1 := mapToListNode(td.ls1)
		ls2 := mapToListNode(td.ls2)
		rs := AddTwoNumbers(ls1, ls2)
		rsList := mapToSlice(rs)
		if !slices.Equal(td.expRes, rsList) {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, rsList, td)
		}
	}
}

func mapToListNode(nums []int) *ListNode {
	var ls *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		ls = &ListNode{
			Val:  nums[i],
			Next: ls,
		}
	}
	return ls
}

func mapToSlice(ls *ListNode) []int {
	var res []int
	for ls != nil {
		res = append(res, ls.Val)
		ls = ls.Next
	}
	return res
}
