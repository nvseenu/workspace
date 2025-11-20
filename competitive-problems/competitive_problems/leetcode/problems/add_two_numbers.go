package problems

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Refer https://leetcode.com/problems/add-two-numbers/description/ for more info.

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs *ListNode

	carryForward := 0
	for l1 != nil || l2 != nil {
		sum := 0
		num1 := 0
		num2 := 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum = num1 + num2 + carryForward
		carryForward = int(sum / 10)
		sum = sum % 10

		rs = &ListNode{
			Val:  sum,
			Next: rs,
		}
	}

	if carryForward > 0 {
		rs = &ListNode{
			Val:  carryForward,
			Next: rs,
		}
	}
	return reverseListNode(rs)
}

func reverseListNode(ls *ListNode) *ListNode {
	var rs *ListNode
	for ls != nil {
		rs = &ListNode{
			Val:  ls.Val,
			Next: rs,
		}
		ls = ls.Next
	}
	return rs
}
