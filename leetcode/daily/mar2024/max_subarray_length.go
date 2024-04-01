package mar2024

// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/
func maxSubarrayLength(nums []int, k int) int {
	l, r, maxLen := 0, 0, 1
	freq := make(map[int]int)
	for r < len(nums) {
		if freq[nums[r]] < k {
			// record result
			maxLen = max(maxLen, r-l+1)
			// move right
			freq[nums[r]]++
			r++
			continue
		}
		if l == r {
			r++
		}
		// move left
		freq[nums[l]]--
		l++
	}
	return maxLen
}
