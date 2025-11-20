package problems

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Refer https://leetcode.com/problems/two-sum/description/ for more info.

func TwoSum(nums []int, target int) []int {

	ns := make(map[int]int)
	for i, v := range nums {
		ns[v] = i
	}

	for i, v := range nums {
		if res := matchTarget(ns, i, v, target); res != nil {
			return res
		}

	}
	return nil

}

func matchTarget(ns map[int]int, originalIdx, originalValue, target int) []int {
	idx, ok := ns[target-originalValue]
	if ok && idx != originalIdx {
		return []int{originalIdx, idx}
	}
	return nil
}
