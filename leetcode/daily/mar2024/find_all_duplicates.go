package mar2024

// https://leetcode.com/problems/find-all-duplicates-in-an-array/
func findDuplicates(nums []int) []int {
	ret := make([]int, 0)
	for i := range nums {
		n := abs(nums[i])
		if nums[n-1] < 0 {
			ret = append(ret, n)
		} else {
			nums[n-1] *= -1
		}
	}
	return ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
