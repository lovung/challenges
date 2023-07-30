package contest356

// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
func countCompleteSubarrays(nums []int) int {
	wholeArrayDistinct := make(map[int]struct{})
	for i := range nums {
		wholeArrayDistinct[nums[i]] = struct{}{}
	}
	distinct := len(wholeArrayDistinct)
	cnt := 0
	for i := 0; i < len(nums); i++ {
		newDistinctMap := make(map[int]struct{})
		for j := i; j < len(nums); j++ {
			newDistinctMap[nums[j]] = struct{}{}
			if len(newDistinctMap) == distinct {
				cnt += len(nums) - j
				break
			}
		}
	}
	return cnt
}
