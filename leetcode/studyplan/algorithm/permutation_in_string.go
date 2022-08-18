package algorithm

const n = int('z' - 'a' + 1)

// Link: https://leetcode.com/problems/permutation-in-string/
// BigO: O(n+m)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	charS1 := make([]int, n)
	charS2 := make([]int, n)
	for i := range s1 {
		charS1[s1[i]-'a']++
		charS2[s2[i]-'a']++
	}

	for i := len(s1); i < len(s2); i++ {
		if checkPermutation(charS1, charS2) {
			return true
		}
		charS2[s2[i-len(s1)]-'a']--
		charS2[s2[i]-'a']++
	}
	if checkPermutation(charS1, charS2) {
		return true
	}
	return false
}

func checkPermutation(s1, s2 []int) bool {
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
