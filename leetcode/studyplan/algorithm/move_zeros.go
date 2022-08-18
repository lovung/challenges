package algorithm

// Link: https://leetcode.com/problems/move-zeroes/
// BigO: O(n)
func moveZeroes(nums []int) {
	newNums := make([]int, len(nums))
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			newNums[j] = nums[i]
			j++
		}
	}
	copy(nums, newNums)
}
