package main

import (
	"fmt"
	"math"
)

func LCSubstring(db [][]int, s1 []byte, startIndex int, endIndex int) int {
	if startIndex > endIndex {
		return 0
	}
	if startIndex == endIndex {
		return 1
	}
	var lcs float64 = 0
	if db[startIndex][endIndex] == 0 {
		if s1[startIndex] == s1[endIndex] {
			remiainingCount := LCSubstring(db, s1, startIndex+1, endIndex-1)
			if remiainingCount == endIndex-startIndex-1 {
				lcs = float64(remiainingCount + 2)
			}
		}
		lcs2 := float64(LCSubstring(db, s1, startIndex+1, endIndex))
		lcs3 := float64(LCSubstring(db, s1, startIndex, endIndex-1))
		db[startIndex][endIndex] = int(math.Max(math.Max(lcs, lcs2), lcs3))
	}
	return db[startIndex][endIndex]
}

func main() {
	str1 := "yabccba"
	strByte := []byte(str1)
	lenStr := len(strByte)
	dp := make([][]int, lenStr)
	for i := 0; i < lenStr; i++ {
		dp[i] = make([]int, lenStr) // initialize a slice of dx unit8 in each of dy slices
	}
	fmt.Println(dp)
	fmt.Println(LCSubstring(dp, strByte, 0, lenStr-1))
}

func maxsum() {

}
