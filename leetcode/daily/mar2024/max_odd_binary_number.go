package mar2024

import "strings"

// https://leetcode.com/problems/maximum-odd-binary-number/?envType=daily-question&envId=2024-03-01
func maximumOddBinaryNumber(s string) string {
	const (
		zero = '0'
		one  = '1'
	)
	cntOne := 0
	for i := range s {
		if s[i] == one {
			cntOne++
		}
	}
	out := strings.Builder{}
	for range cntOne - 1 {
		out.WriteByte(one)
	}
	for range len(s) - cntOne {
		out.WriteByte(zero)
	}

	out.WriteByte(one)
	return out.String()
}
