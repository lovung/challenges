package jun2024

import "math"

// https://leetcode.com/problems/sum-of-square-numbers/description/
func judgeSquareSum(c int) bool {
	sc := int(math.Sqrt(float64(c)))
	for a := 0; a <= sc; a++ {
		b2 := c - a*a
		b := int(math.Sqrt(float64(b2)))
		if b*b == b2 {
			return true
		}
	}
	return false
}
