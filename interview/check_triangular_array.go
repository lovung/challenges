package interview

func checkTriangularArray(nums []int) bool {
	var (
		direction = 1
		l, r      = -1, 0
		N         = len(nums)
	)
	for r < N {
		if direction == 1 && r > 0 {
			if nums[r-1] < nums[r] {
				r++
				continue
			} else if nums[r-1] == nums[r] {
				return false
			}
			direction--
			l = r - 1
		} else if direction == 0 {
			if nums[r-1] < nums[r] {
				return false
			}
		}
		for l >= 0 {
			if nums[l] > nums[r] {
				l--
			} else if nums[l] == nums[r] {
				return false
			} else {
				break
			}
		}
		r++
	}
	return direction == 0
}
