package may2024

import "github.com/lovung/ds/maths"

func equalSubstring(s string, t string, maxCost int) int {
	costs := calcCosts(s, t)
	res := 0
	for i := range costs {
		for j := i + res + 1; j < len(costs); j++ {
			if costs[j]-costs[i] <= maxCost {
				res = max(res, j-i)
			} else {
				break
			}
		}
	}
	return res
}

func calcCosts(s, t string) []int {
	prefixSumCosts := make([]int, len(s)+1)
	sum := 0
	for i := range s {
		sum += maths.ABS(int(s[i]) - int(t[i]))
		prefixSumCosts[i+1] = sum
	}
	return prefixSumCosts
}
