package oct2022

import (
	"strconv"
	"strings"
)

// Link: https://leetcode.com/problems/minimize-result-by-adding-parentheses-to-expression/
func minimizeResult(expression string) string {
	// expression: a+b
	// prefix_a * (suffix_a + prefix_b) * suffix_b
	lenA := strings.Index(expression, "+")
	lenB := len(expression) - 1 - lenA

	a, err := strconv.Atoi(expression[:lenA])
	if err != nil {
		return ""
	}
	b, err := strconv.Atoi(expression[lenA+1:])
	if err != nil {
		return ""
	}
	minRes := a + b
	minI, minJ := 0, lenB
	for i := 0; i < lenA; i++ {
		for j := 1; j <= lenB; j++ {
			prefixA, suffixA := breakBy(a, lenA, i)
			prefixB, suffixB := breakBy(b, lenB, j)
			val := prefixA * (suffixA + prefixB) * suffixB
			if val < minRes {
				minRes = val
				minI, minJ = i, j
			}
		}
	}
	return expression[:minI] + "(" + expression[minI:lenA+1+minJ] + ")" + expression[lenA+1+minJ:]
}

func tenPow(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}

// return prefix and suffix in integer
func breakBy(num, lenNum, at int) (int, int) {
	if at == 0 {
		return 1, num
	}
	if at == lenNum {
		return num, 1
	}
	return num / tenPow(lenNum-at), num % tenPow(lenNum-at)
}
