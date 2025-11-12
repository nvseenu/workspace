package sorting

func MergeSort[E any](elements []E, compare func(a, b E) int) {
	mergeSortInternal(elements, 0, len(elements)-1, compare)
}

func mergeSortInternal[E any](arr []E, si int, ei int, compareFn func(a, b E) int) {
	// Stop if we can't split the array further
	if ei-si == 0 {
		return
	}
	mid := si + int((ei-si)/2)
	mergeSortInternal(arr, si, mid, compareFn)
	mergeSortInternal(arr, mid+1, ei, compareFn)
	merge(arr, si, mid, ei, compareFn)
}

func copySubArray[E any](arr []E, si int, ei int) []E {
	subArrLen := ei - si + 1
	subArr := make([]E, subArrLen)
	subArrIdx := 0
	for i := si; i <= ei; i++ {
		subArr[subArrIdx] = arr[i]
		subArrIdx++
	}
	return subArr
}

func merge[E any](arr []E, si int, mid int, ei int, compareFn func(a, b E) int) {
	leftArr := copySubArray(arr, si, mid)
	rightArr := copySubArray(arr, mid+1, ei)

	li := 0
	ri := 0
	ai := si
	for li < len(leftArr) && ri < len(rightArr) {

		if compareFn(leftArr[li], rightArr[ri]) == -1 {
			arr[ai] = rightArr[ri]
			ri++
		} else {
			arr[ai] = leftArr[li]
			li++
		}
		ai++
	}

	// Copy remaining elements from left array if any
	for i := li; i < len(leftArr); i++ {
		arr[ai] = leftArr[i]
		ai++
	}

	// Copy remaining elements from right array if any
	for i := ri; i < len(rightArr); i++ {
		arr[ai] = rightArr[i]
		ai++
	}
}
