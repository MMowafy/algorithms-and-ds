package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{10, 5, 30, 15, 50, 6}))
}

func bubbleSort(arr []int) []int {
	startIndex := 0
	endIndex := len(arr) - 1
	for startIndex < endIndex {
		for i := startIndex; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
		endIndex--
	}
	return arr
}
