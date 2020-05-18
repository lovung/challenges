package solution

import "math"

// TapeEquilibrium problem
// @Link: https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/
func TapeEquilibrium(A []int) int {
	length := len(A)
	sumA := make([]int, length)
	sum := 0
	for i, e := range A {
		sum += e
		sumA[i] = sum
	}
	var min int = math.MaxInt64
	for i := 0; i < length-1; i++ {
		tmp := sum - 2*sumA[i]
		if tmp < 0 {
			tmp = -tmp
		}
		if tmp < min {
			min = tmp
		}
	}
	return min
}
