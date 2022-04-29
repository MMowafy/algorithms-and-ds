package main

import "fmt"

func main() {
	fmt.Println(insertionSort([]int{2, 0, 2, 1, 1, 0}))
}

func insertionSort(items []int) []int {
	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j-1] > items[j] {
				temp := items[j]
				items[j] = items[j-1]
				items[j-1] = temp
			}
		}
	}
	return items
}
