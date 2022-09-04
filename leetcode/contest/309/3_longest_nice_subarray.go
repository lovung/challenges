package contest

// Link: https://leetcode.com/problems/longest-nice-subarray/
func longestNiceSubarray(nums []int) int {
	n := len(nums)
	l, r := 0, 0
	curMask := 0
	max := 1
	for r < n {
		for (l < r) && (curMask&nums[r] != 0) {
			curMask ^= nums[l]
			l++
		}
		curMask |= nums[r]
		r++
		if max < r-l {
			max = r - l
		}
	}
	return max
}
