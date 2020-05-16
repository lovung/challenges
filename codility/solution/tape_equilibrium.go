package solution

// you can also use imports, for example:
// import "fmt"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// TapeEquilibrium problem
// @Link: https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/
func TapeEquilibrium(A []int) int {
	// write your code in Go 1.4
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
