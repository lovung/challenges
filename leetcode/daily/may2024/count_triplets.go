package may2024

// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
func countTriplets(arr []int) int {
	xor := 0
	prevIdxes := make(map[int][]int)
	prevIdxes[0] = append(prevIdxes[0], -1)
	res := 0
	for k, a := range arr {
		xor ^= a
		for _, i := range prevIdxes[xor] {
			res += k - i - 1
		}
		prevIdxes[xor] = append(prevIdxes[xor], k)
	}
	return res
}
