package algorithm

// Link: https://leetcode.com/problems/search-insert-position/
// BigO: O(log n)
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for {
		m := (l + r) / 2
		switch {
		// stop conditions
		case nums[m] == target:
			return m
		case l == m && nums[l] > target:
			return l
		case l == m && nums[l] < target:
			return r
		default:
			// next actions
			if nums[m] > target {
				r = m
			} else {
				l = m
			}
		}

	}
}
