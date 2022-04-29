package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(bucketSort([]int{10, 90, 80, 100, 60, 50, 30, 40, 20}))
}

func bucketSort(items []int) []int {
	max := math.MinInt32
	numberOfBuckets := math.Ceil(math.Sqrt(float64(len(items))))
	bucketsArr := make([][]int, int(numberOfBuckets))
	for _, value := range items {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
	for _, value := range items {
		valueBucket := int(math.Ceil(float64(value) * numberOfBuckets / float64(max)))
		fmt.Println(valueBucket)
		bucketsArr[valueBucket-1] = append(bucketsArr[valueBucket-1], value)
	}

	for _, bucket := range bucketsArr {
		sort.Ints(bucket)
	}
	items = []int{}
	for _, bucket := range bucketsArr {
		items = append(items, bucket...)
	}
	return items
}
