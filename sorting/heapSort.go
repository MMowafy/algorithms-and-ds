package main

import (
	"data-structure-and-algorithms/heap"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(20)
	fmt.Println("slice ", slice)
	heapSort(slice)
}

func heapSort(arr []int) {
	heapArr := heap.NewMinHeap(len(arr))
	for i := 0; i < len(arr); i++ { //Insert in Heap
		heapArr.Insert(arr[i])
	}

	for i := 0; i < len(arr); i++ { //Extract from Heap and insert sorted data in current Array
		arr[i] = heapArr.Remove()
	}
	fmt.Println(arr)
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999)
	}
	return slice
}
