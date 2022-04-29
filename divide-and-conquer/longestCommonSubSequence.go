package main

import (
	"fmt"
	"math"
)

func LCS(s1 []byte, s2 []byte, s1Index int, S2Index int, db [][]int) int {
	if s1Index >= len(s1) || S2Index >= len(s2) {
		return 0
	}
	var lcs float64 = 0
	if db[s1Index][S2Index] != 0 {
		return db[s1Index][S2Index]
	}
	if s1[s1Index] == s2[S2Index] {
		lcs = 1 + float64(LCS(s1, s2, s1Index+1, S2Index+1, db))
	}
	lcs2 := float64(LCS(s1, s2, s1Index+1, S2Index, db))
	lcs3 := float64(LCS(s1, s2, s1Index, S2Index+1, db))
	db[s1Index][S2Index] = int(math.Max(math.Max(lcs, lcs2), lcs3))
	return db[s1Index][S2Index]
}

func main() {
	str1 := "HARRY"
	str2 := "SALLY"
	strByte := []byte(str1)
	lenStr := len(strByte)
	str2Byte := []byte(str1)
	len2Str := len(str2Byte)
	dp := make([][]int, lenStr)
	for i := 0; i < len2Str; i++ {
		dp[i] = make([]int, len2Str) // initialize a slice of dx unit8 in each of dy slices
	}
	fmt.Println(LCS([]byte(str1), []byte(str2), 0, 0, dp))
}
