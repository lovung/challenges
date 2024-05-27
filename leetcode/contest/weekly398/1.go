package weekly398

func isArraySpecial1(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if i&1 == 0 {
			if (nums[i] & 1) != (nums[0] & 1) {
				return false
			}
		} else {
			if (nums[i] & 1) == (nums[0] & 1) {
				return false
			}
		}
	}
	return true
}
