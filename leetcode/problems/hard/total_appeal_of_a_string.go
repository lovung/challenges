package hard

import (
	"fmt"
)

// Link: https://leetcode.com/problems/total-appeal-of-a-string/
const (
	a    = 'a'
	atoz = 26
)

// O(26*n^2) solution
func appealSum(s string) int64 {
	l := len(s)
	sum := l
	currentMark := [atoz]int{}
	for i := range s {
		currentMark[s[i]-a]++
		// skip first charactor because we have process as assigning 1 above
		if i == 0 {
			continue
		}
		newMark := [atoz]int{}
		copy(newMark[:], currentMark[:])
		sum += countAppeal(newMark[:])
		for j := 1; j < l-i; j++ {
			newMark[s[j-1]-a]--
			newMark[s[i+j]-a]++
			sum += countAppeal(newMark[:])
		}
	}
	return int64(sum)
}

// O(26) ~ O(1)
func countAppeal(mark []int) int {
	count := 0
	for i := range mark {
		if mark[i] > 0 {
			count++
		} else if mark[i] < 0 {
			panic("should not happen")
		}
	}
	return count
}

func printMark(mark []int) {
	for i := range mark {
		if mark[i] > 0 {
			fmt.Printf(string(rune(i + a)))
		}
	}
	fmt.Println()
}

func sum2DSlice(mat [][]int) int {
	sum := 0
	for i := range mat {
		for j := range mat[i] {
			sum += mat[i][j]
		}
	}
	return sum
}
