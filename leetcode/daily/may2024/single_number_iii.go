package may2024

func singleNumber(nums []int) []int {
	exists := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := exists[n]; ok {
			delete(exists, n)
		} else {
			exists[n] = struct{}{}
		}
	}
	res := make([]int, 0, 2)
	for k := range exists {
		res = append(res, k)
	}
	return res
}

func singleNumber2(nums []int) []int {
	xs := 0
	for _, x := range nums {
		xs ^= x
	}
	// The expression -ans is the two's complement of ans,
	// which is equivalent to ~ans + 1. The bitwise AND of
	// ans and -ans isolates the rightmost set bit.
	lb := xs & -xs
	a := 0
	for _, x := range nums {
		if x&lb != 0 {
			a ^= x
		}
	}
	b := xs ^ a
	return []int{a, b}
}
