package topinterview150

// Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	return _removeDuplicates(nums, 1)
}

// Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
func removeDuplicatesII(nums []int) int {
	return _removeDuplicates(nums, 2)
}

func _removeDuplicates(nums []int, maxE int) int {
	k := len(nums)
	for l := 0; l < k; {
		r := l + 1
		for ; r <= k; r++ {
			if r < k && nums[l] == nums[r] {
				continue
			}
			// Need to shift
			shift := r - (l + maxE)
			if shift > 0 {
				copy(nums[l+maxE:], nums[r:k])
				k -= shift
			}
			break
		}

		l += min(maxE, max(r-(l+maxE), 1))
	}
	return k
}
