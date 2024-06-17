package biweekly132

// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-ii/description/
func maximumLength(nums []int, k int) int {
	dp := make([]map[int]int, k+1)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	res := make([]int, k+1)
	for _, n := range nums {
		for i := k; i >= 0; i-- {
			v := dp[i][n] + 1
			if i > 0 {
				v = max(v, res[i-1]+1)
			}
			dp[i][n] = v
			res[i] = max(res[i], v)
		}
	}
	return res[k]
}
