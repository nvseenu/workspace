package sorting

import (
	"math/rand"
	"reflect"
	"testing"
)

const SortTypeMerge = "MERGE_SORT"

func compareInt(a, b int) int {
	if a > b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

type sortTestData struct {
	Case     string
	Input    []int
	Expected []int
}

var lengthyNumbers = prepareLengthyNumberSlice(1000)
var lengthyNumbersShuffled = shuffle(lengthyNumbers)

func shuffle(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	rand.Shuffle(1000, func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}

func prepareLengthyNumberSlice(n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

var testData = []sortTestData{
	{
		Case:     "All numbers in reverse",
		Input:    []int{5, 4, 3, 2, 1},
		Expected: []int{1, 2, 3, 4, 5},
	},
	{
		Case:     "Unsorted",
		Input:    []int{5, 4, 2, 1, 3},
		Expected: []int{1, 2, 3, 4, 5},
	},

	{
		Case:     "Already sorted",
		Input:    []int{1, 2, 3, 4, 5},
		Expected: []int{1, 2, 3, 4, 5},
	},
	{
		Case:     "Lenghty numbers",
		Input:    lengthyNumbersShuffled,
		Expected: lengthyNumbers,
	},
}

var testSortTypes = []string{SortTypeMerge}

func TestSort(t *testing.T) {

	for _, sortType := range testSortTypes {
		for _, td := range testData {
			input := make([]int, len(td.Input))
			copy(input, td.Input)

			if sortType == SortTypeMerge {
				MergeSort(input, compareInt)
			}

			if !reflect.DeepEqual(td.Expected, input) {
				t.Errorf("Sorting Func: %v, Case: %v , input: %v => Expected %v , but got %v", "MergeSort", td.Case, td.Input, td.Expected, input)
			}
		}
	}

}
