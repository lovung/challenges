package contest

import "github.com/lovung/ds/maths"

// https://leetcode.com/problems/max-pair-sum-in-an-array/
func maxSum(nums []int) int {
	mapMaxDigit := make(map[int]int)
	ret := -1
	for _, n := range nums {
		maxD := maxDigit(n)
		if mapMaxDigit[maxD] == 0 {
			mapMaxDigit[maxD] = n
			continue
		}
		ret = maths.Max(ret, mapMaxDigit[maxD]+n)
		mapMaxDigit[maxD] = maths.Max(mapMaxDigit[maxD], n)
	}
	return ret
}

func maxDigit(num int) int {
	maxD := 0
	for num > 0 {
		maxD = maths.Max(maxD, num%10)
		num /= 10
	}
	return maxD
}
