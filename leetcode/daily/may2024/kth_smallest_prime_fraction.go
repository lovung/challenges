package may2024

import "sort"

// https://leetcode.com/problems/k-th-smallest-prime-fraction
func kthSmallestPrimeFraction(arr []int, k int) []int {
	type fraction struct {
		a, b int
	}
	fractions := make([]fraction, 0, len(arr)*len(arr))
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			fractions = append(fractions, fraction{arr[i], arr[j]})
		}
	}
	sort.Slice(fractions, func(i, j int) bool {
		return fractions[i].a*fractions[j].b < fractions[i].b*fractions[j].a
	})
	return []int{fractions[k-1].a, fractions[k-1].b}
}
