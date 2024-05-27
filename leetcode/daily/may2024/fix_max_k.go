package may2024

func findMaxK(nums []int) int {
	exist := make(map[int]bool)
	for i := range nums {
		if nums[i] < 0 {
			exist[nums[i]] = true
		}
	}
	res := -1
	for i := range nums {
		if nums[i] > 0 && exist[-nums[i]] {
			res = max(res, nums[i])
		}
	}
	return res
}
