package main

import (
	"dsa_go/sorting"
	"fmt"
)

func main() {

	arr := []int{3, 2, 1}
	sorting.MergeSort[int](arr, func(a, b int) int {
		return -1
	})

	fmt.Println("sort ", arr)
}
