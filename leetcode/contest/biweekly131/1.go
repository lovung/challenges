package biweekly131

// https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/submissions/
func duplicateNumbersXOR(nums []int) int {
	res := 0
	exist := make(map[int]bool)
	for i := range nums {
		if exist[nums[i]] {
			res ^= nums[i]
		}
		exist[nums[i]] = true
	}
	return res
}
