package algorithm

func rotate(nums []int, k int) {
	n := len(nums)
	k -= (k / n) * n
	newNums := make([]int, n)
	copy(newNums, nums[n-k:])
	copy(newNums[k:], nums[:n-k])
	copy(nums, newNums)
}
