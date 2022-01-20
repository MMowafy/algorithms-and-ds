package main

import "fmt"

func NumberFactor(num int, arr []int) int {
	if num == 0 || num == 1 || num == 2 {
		return 1
	}
	if num == 3 {
		return 2
	}

	num1 := NumberFactor(num-1, arr)
	num2 := NumberFactor(num-3, arr)
	num3 := NumberFactor(num-4, arr)
	return num1 + num2 + num3
}

func main() {
	fmt.Println(NumberFactor(5, []int{1, 3, 4}))
}
