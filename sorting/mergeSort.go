package main

import (
	"fmt"
)

func main() {
	fmt.Println(mergeSortDC([]int{10, 90, 80, 70, 100, 110, 50, 5}))
}

func mergeSortDC(items []int) []int {
	if len(items) == 1 {
		return items
	}
	middle := int(len(items) / 2)
	left := make([]int, middle)
	right := make([]int, len(items)-middle)
	for i := 0; i < len(items); i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	sortedLeft := mergeSortDC(left)
	sortedRight := mergeSortDC(right)
	return merge(sortedLeft, sortedRight)
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0

	// compare two lists
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
			i++
		} else {
			result[i] = right[0]
			right = right[1:]
			i++
		}
	}
	// add remianing from left if exists
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	// add remianing from right if exists
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
