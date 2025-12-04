package common

func IsArraySorted(nums []int) bool {
	return isArraySortedInternal(nums, len(nums)-1)
}

func isArraySortedInternal(nums []int, index int) bool {
	if index == 0 {
		return true
	}

	if nums[index] > nums[index-1] {
		return isArraySortedInternal(nums, index-1)
	} else {
		return false
	}

}
