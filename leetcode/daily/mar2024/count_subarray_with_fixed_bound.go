package mar2024

// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/description/
func countSubarraysWithFixedBounds(nums []int, minK int, maxK int) int64 {
	a := countSubarraysWithBounds(nums, minK, maxK)
	b := countSubarraysWithBounds(nums, minK+1, maxK)
	c := countSubarraysWithBounds(nums, minK, maxK-1)
	d := countSubarraysWithBounds(nums, minK+1, maxK-1)
	return a - b - c + d
}

func countSubarraysWithBounds(nums []int, minK int, maxK int) int64 {
	res := int64(0)
	for l, r := 0, 0; r < len(nums); r++ {
		if !inBound(nums[r], minK, maxK) {
			l = r + 1
			continue
		}
		res += int64(r) - int64(l) + 1
	}
	return res
}

func inBound(n, minK, maxK int) bool {
	return n >= minK && n <= maxK
}
