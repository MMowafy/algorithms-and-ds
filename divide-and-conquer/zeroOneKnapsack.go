package main

import (
	"fmt"
	"math"
)

func ZeroOneKnapsack(weights []float64, profits []float64, capacity float64, index int) float64 {
	if index >= len(profits) || capacity <= 0 {
		return 0
	}
	var p1 float64 = 0
	if capacity >= weights[index] {
		p1 = ZeroOneKnapsack(weights, profits, capacity-weights[index], index+1) + profits[index]
	}
	p2 := ZeroOneKnapsack(weights, profits, capacity, index+1)
	return math.Max(p1, p2)
}

func main() {
	fmt.Println(ZeroOneKnapsack([]float64{3, 1, 5, 2}, []float64{31, 26, 72, 17}, 7, 0))
}
