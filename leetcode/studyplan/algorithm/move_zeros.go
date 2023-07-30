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

func moveZeroes2(nums []int) {
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			continue
		}
		nums[l] = nums[r]
		l++
	}
	for ; l < len(nums); l++ {
		nums[l] = 0
	}
}
