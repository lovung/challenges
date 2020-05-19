package solution

// PermMissingElem solution
// link: https://app.codility.com/demo/results/trainingXT35B7-C96/
func PermMissingElem(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	sum := (N + 1) * (N + 2) / 2
	count := 0
	for _, e := range A {
		count += e
	}
	return sum - count
}
