package main

import (
	"fmt"
	"math"
)

func LCS(s1 []byte, s2 []byte, s1Index int, S2Index int) int {
	if s1Index >= len(s1) || S2Index >= len(s2) {
		return 0
	}
	var lcs float64 = 0
	if s1[s1Index] == s2[S2Index] {
		lcs = 1 + float64(LCS(s1, s2, s1Index+1, S2Index+1))
	}
	lcs2 := float64(LCS(s1, s2, s1Index+1, S2Index))
	lcs3 := float64(LCS(s1, s2, s1Index, S2Index+1))
	return int(math.Max(math.Max(lcs, lcs2), lcs3))
}

func main() {
	str1 := "elephant"
	str2 := "eretpat"
	fmt.Println(LCS([]byte(str1), []byte(str2), 0, 0))
}
