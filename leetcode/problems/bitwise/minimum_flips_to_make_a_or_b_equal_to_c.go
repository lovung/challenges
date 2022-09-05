package bitwise

// Link: https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
func minFlips(a int, b int, c int) int {
	mask := 1
	cnt := 0
	for mask != 0 {
		if c&mask != 0 {
			// if bit in c is 1, need as least (min) 1 bit in a or b is 1
			if (a|b)&mask == 0 {
				cnt++
			}
		} else {
			// if bit in c is 0, need to flip all 1-bit in a and b to 0
			if a&mask != 0 {
				cnt++
			}
			if b&mask != 0 {
				cnt++
			}
		}
		mask <<= 1
	}
	return cnt
}
