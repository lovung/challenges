package jul2023

// https://leetcode.com/problems/predict-the-winner/editorial/
// Solution 1:
func PredictTheWinner(nums []int) bool {
	return maxDiff(nums, 0, len(nums)-1) >= 0
}

func maxDiff(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	return max(
		nums[l]-maxDiff(nums, l+1, r),
		nums[r]-maxDiff(nums, l, r-1),
	)
}

// Solution 2:
func PredictTheWinner2(nums []int) bool {
	cache := make([][]int, len(nums))
	for i := range cache {
		cache[i] = make([]int, len(nums))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return maxDiff2(nums, 0, len(nums)-1, cache) >= 0
}

func maxDiff2(nums []int, l, r int, cache [][]int) int {
	if cache[l][r] != -1 {
		return cache[l][r]
	}
	if l == r {
		return nums[l]
	}
	cache[l][r] = max(
		nums[l]-maxDiff2(nums, l+1, r, cache),
		nums[r]-maxDiff2(nums, l, r-1, cache),
	)
	return cache[l][r]
}

// Solotion 3:
func PredictTheWinner3(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for diff := 1; diff < n; diff++ {
		for l := 0; l < n-diff; l++ {
			r := l + diff
			dp[l][r] = max(
				nums[l]-dp[l+1][r],
				nums[r]-dp[l][r-1])
		}
	}
	return dp[0][n-1] >= 0
}

// Solution 4:
func PredictTheWinner4(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)

	copy(dp, nums)
	for diff := 1; diff < n; diff++ {
		for l := 0; l < n-diff; l++ {
			r := l + diff
			dp[l] = max(
				nums[l]-dp[l+1],
				nums[r]-dp[l])
		}
	}
	return dp[0] >= 0
}
