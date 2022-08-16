package algorithm

import "sort"

// Link: https://leetcode.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	squares := make([]int, len(nums))
	for i := range nums {
		squares[i] = nums[i] * nums[i]
	}
	sort.Ints(squares)
	return squares
}

func sortedSquares2(nums []int) []int {
	l, r := 0, len(nums)-1
	squares := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[l] < 0 && -nums[l] > nums[r] {
			squares[i] = nums[l] * nums[l]
			l++
		} else {
			squares[i] = nums[r] * nums[r]
			r--
		}
	}
	return squares
}
