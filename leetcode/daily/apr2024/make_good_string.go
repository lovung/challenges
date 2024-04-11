package apr2024

// https://leetcode.com/problems/make-the-string-great/description/
func makeGood(s string) string {
	picked := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(picked) > 0 && isBad(picked[len(picked)-1], s[i]) {
			picked = picked[:len(picked)-1]
		} else {
			picked = append(picked, s[i])
		}
	}
	return string(picked)
}

func isBad(a, b byte) bool {
	const delta = 'a' - 'A'
	return a-b == delta || b-a == delta
}
