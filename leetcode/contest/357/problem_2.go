package contest

import "github.com/lovung/ds/maths"

// Recursive solution
func canSplitArray(nums []int, m int) bool {
	sum := maths.Sum(nums...)

	return recursiveCanSplitArray(nums, m, sum)
}

func recursiveCanSplitArray(nums []int, m int, sum int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}
	ret := false
	if sum-nums[0] >= m {
		ret = ret || recursiveCanSplitArray(nums[1:], m, sum-nums[0])
	}
	if sum-nums[n-1] >= m {
		ret = ret || recursiveCanSplitArray(nums[:n-1], m, sum-nums[n-1])
	}
	return ret
}

// DP
func canSplitArray2(nums []int, m int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// col is k+i
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			j := i + k
			if j >= n {
				break
			}
			switch k {
			case 0:
				dp[i][j] = nums[j]
			case 1:
				dp[i][j] = dp[i][j-1] + nums[j]
			default:
				if dp[i][j-1] < m && dp[i+1][j] < m {
					dp[i][j] = 0
				} else {
					dp[i][j] = maths.Max(dp[i][j-1], dp[i+1][j]) + nums[j]
				}
			}
		}
	}
	return dp[0][n-1] > 0
}

// Smartest solution
func canSplitArray3(nums []int, m int) bool {
	if len(nums) <= 2 {
		return true
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] >= m {
			return true
		}
	}
	return false
}
