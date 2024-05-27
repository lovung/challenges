package biweekly131

// https://leetcode.com/problems/find-occurrences-of-an-element-in-an-array/description/
func occurrencesOfElement(nums []int, queries []int, x int) []int {
	idx := make([]int, 0, len(nums))
	for i, num := range nums {
		if num == x {
			idx = append(idx, i)
		}
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		if q > len(idx) {
			res[i] = -1
		} else {
			res[i] = idx[q-1]
		}
	}
	return res
}
