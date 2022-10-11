package oct2022

// Link: https://leetcode.com/problems/increasing-triplet-subsequence/
func increasingTriplet(nums []int) bool {
	dp := make([]int, len(nums))
	size := 0
	for i := range nums {
		l, r := 0, size
		for l < r {
			m := (l + r) / 2
			if dp[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		dp[l] = nums[i]
		if l == size {
			size++
			if size >= 3 {
				break
			}
		}
	}
	return size >= 3
}
