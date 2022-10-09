package contest

// Link: https://leetcode.com/problems/find-the-original-array-of-prefix-xor
func findArray(pref []int) []int {
	res := make([]int, len(pref))
	prevPref := 0
	for i := range pref {
		res[i] = pref[i] ^ prevPref
		prevPref = pref[i]
	}
	return res
}
