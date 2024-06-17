package jun2024

import "sort"

// https://leetcode.com/problems/height-checker/
func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	res := 0
	sort.Ints(expected)
	for i := range heights {
		if heights[i] != expected[i] {
			res++
		}
	}
	return res
}
