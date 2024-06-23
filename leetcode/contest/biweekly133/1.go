package biweekly133

// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/description/
func minimumOperations(nums []int) int {
	res := 0
	for i := range nums {
		if nums[i]%3 != 0 {
			res++
		}
	}
	return res
}
