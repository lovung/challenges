package may2024

// https://leetcode.com/problems/subsets
func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, 1<<5)
	for i := range 1 << n {
		res = append(res, subsetWithBitmask(nums, i))
	}
	return res
}

func subsetWithBitmask(nums []int, bitmask int) []int {
	subset := make([]int, 0, len(nums))
	bit := 1 << (len(nums) - 1)
	for i := range nums {
		if bitmask&bit != 0 {
			subset = append(subset, nums[i])
		}
		bit >>= 1
	}
	return subset
}
