package weekly397

import "math"

// https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/description/
func maximumEnergy(energy []int, k int) int {
	res := math.MinInt
	for i := len(energy) - 1; i >= len(energy)-k; i-- {
		sum := 0
		for j := i; j >= 0; j -= k {
			sum += energy[j]
			res = max(res, sum)
		}
	}
	return res
}
