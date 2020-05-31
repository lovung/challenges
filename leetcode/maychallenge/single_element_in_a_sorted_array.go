package maychallenge

// Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3327/
func singleNonDuplicate(nums []int) int {
	length := len(nums)
	var mid int
	for l, r := 0, length-1; l <= r; {
		mid = (l + r) / 2
		if l == r {
			break
		}
		if (mid-l)%2 == 0 {
			if nums[mid] == nums[mid+1] {
				l = mid
				continue
			}
			if nums[mid] == nums[mid-1] {
				r = mid
				continue
			}
			break
		}
		if nums[mid] == nums[mid+1] {
			r = mid - 1
			continue
		}
		if nums[mid] == nums[mid-1] {
			l = mid + 1
			continue
		}
	}
	return nums[mid]
}
