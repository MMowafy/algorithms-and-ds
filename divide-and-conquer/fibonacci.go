package main

import "fmt"

func Fibonacci(num int) int {
	// base case
	if num <= 1 {
		return num
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}

func main() {
	fmt.Println(Fibonacci(9))
}
