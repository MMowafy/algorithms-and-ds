package main

import (
	"fmt"
	"math"
)

func LCPS(s1 []byte, startIndex int, endIndex int) int {
	if startIndex > endIndex {
		return 0
	}
	if startIndex == endIndex {
		return 1
	}
	var lcs float64 = 0
	if s1[startIndex] == s1[endIndex] {
		lcs = 2 + float64(LCPS(s1, startIndex+1, endIndex-1))
	}
	lcs2 := float64(LCPS(s1, startIndex+1, endIndex))
	lcs3 := float64(LCPS(s1, startIndex, endIndex-1))
	return int(math.Max(math.Max(lcs, lcs2), lcs3))
}

func main() {
	str1 := "mamdrdm"
	fmt.Println(LCPS([]byte(str1), 0, 6))
}
