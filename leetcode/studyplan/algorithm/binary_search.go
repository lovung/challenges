package algorithm

func search(nums []int, target int) int {
	l, r := 0, len(nums)
	mid := (l + r) / 2
	for l != mid {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}
	if nums[mid] == target {
		return mid
	}
	return -1
}
