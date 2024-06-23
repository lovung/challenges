package biweekly133

func minOperations2(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			res++
		}
	}
	if nums[0] == 0 {
		res++
	}
	return res
}
