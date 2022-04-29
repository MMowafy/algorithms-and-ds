package main

import (
	"fmt"
)

func main() {
	slice := generateSlice(20)
	fmt.Println(quickSort(slice))
}

//func quickSort(a []int) []int {
//	if len(a) < 2 {
//		return a
//	}
//
//	left, right := 0, len(a)-1
//
//	pivot := len(a) - 1
//
//	a[pivot], a[right] = a[right], a[pivot]
//
//	for i, _ := range a {
//		if a[i] < a[right] {
//			a[left], a[i] = a[i], a[left]
//			left++
//		}
//	}
//
//	a[left], a[right] = a[right], a[left]
//
//	quickSort(a[:left])
//	quickSort(a[left+1:])
//
//	return a
//}

func quickSort(items []int) []int {
	pivotIndex := partition(items)
	if len(items) <= 1 {
		return items
	}
	leftArr := make([]int, pivotIndex)
	rightArr := make([]int, len(items)-pivotIndex-1)
	leftArr = items[:pivotIndex]
	rightArr = items[pivotIndex:]
	sortedLeftArr := quickSort(leftArr)
	SortedRightArr := quickSort(rightArr)
	result := sortedLeftArr
	result = append(result, SortedRightArr...)
	return result
}

func partition(items []int) int {
	pivot := len(items) - 1
	i := -1
	for j, _ := range items {
		if items[pivot] >= items[j] {
			i++
			temp := items[i]
			items[i] = items[j]
			items[j] = temp
		}
	}
	return i
}

//func generateSlice(size int) []int {
//
//	slice := make([]int, size, size)
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < size; i++ {
//		slice[i] = rand.Intn(999)
//	}
//	return slice
//}
