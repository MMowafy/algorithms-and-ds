package main

import (
	"fmt"
)

func main() {
	fmt.Println(Selectionsort([]int{10, 5, 30, 15, 50, 6}))
}

func Selectionsort(items []int) []int {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i + 1; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		temp := items[i]
		items[i] = items[minIdx]
		items[minIdx] = temp
	}
	return items
}
