package mar2024

// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/description/?envType=daily-question&envId=2024-03-05
func minimumLength(s string) int {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			break
		}
		c := s[l]
		for ; l <= r && s[l] == c; l++ {
		}
		for ; l <= r && s[r] == c; r-- {
		}
	}
	return r - l + 1
}
