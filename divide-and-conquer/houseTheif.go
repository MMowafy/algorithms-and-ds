package main

import (
	"fmt"
	"math"
)

func HouseThiefMaxMoney(arr []int, index int) float64 {
	if index >= len(arr) {
		return 0
	}
	sum1 := float64(arr[index]) + HouseThiefMaxMoney(arr, index+2)
	sum2 := HouseThiefMaxMoney(arr, index+1)
	return math.Max(sum1, sum2)
}

func main() {
	fmt.Println(HouseThiefMaxMoney([]int{6, 7, 1, 30, 8, 2, 4}, 0))
}
