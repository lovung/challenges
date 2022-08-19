package problems

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3321/
 */

func majorityElement(nums []int) int {
	// Solution MM
	// mm := make(map[int]int)
	// length := len(nums) >> 1
	// for _, e := range nums {
	//     mm[e]++
	//     if mm[e] > length {
	//         return e
	//     }
	// }
	// return 0

	// Solution sort
	// sort.Ints(nums)
	// return nums[len(nums) >> 1]

	// Solution up/down
	count := 0
	var res int
	for _, e := range nums {
		if count == 0 {
			res, count = e, 1
		} else {
			if e == res {
				count++
			} else {
				count--
			}
		}
	}
	return res
}
