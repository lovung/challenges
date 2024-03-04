package mar2024

// https://leetcode.com/problems/find-the-duplicate-number/
func findDuplicate1(nums []int) int {
	exist := make([]bool, len(nums)+1)
	for i := range nums {
		if exist[nums[i]] {
			return nums[i]
		}
		exist[nums[i]] = true
	}
	return 0
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); {
		n := nums[i]
		if n == i+1 {
			i++
		} else if n == nums[n-1] {
			return n
		} else {
			nums[i], nums[n-1] = nums[n-1], n
		}
	}
	return 0
}
