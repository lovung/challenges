package contest

// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/
func countPairs(nums []int, target int) int {
	n := len(nums)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] < target {
				cnt++
			}
		}
	}
	return cnt
}
