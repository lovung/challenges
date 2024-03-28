package grind75

// https://leetcode.com/problems/house-robber/
func rob(nums []int) int {
	take, dont := 0, 0
	for i := range nums {
		take, dont = nums[i]+dont, max(take, dont)
	}
	return max(take, dont)
}

//         2 7 9  3  1
//  take 0 2 7 11 10 12
//  dont 0 0 2 7  11 11
