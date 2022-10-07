package grind75

// Link: https://leetcode.com/problems/sort-colors/
func sortColors(nums []int) {
	redCnt, whiteCnt, blueCnt := 0, 0, 0
	for i := range nums {
		switch nums[i] {
		case 0:
			redCnt++
		case 1:
			whiteCnt++
		case 2:
			blueCnt++
		}
	}
	for i := range nums {
		switch {
		case i < redCnt:
			nums[i] = 0
		case i < redCnt+whiteCnt:
			nums[i] = 1
		default:
			nums[i] = 2
		}
	}
}
