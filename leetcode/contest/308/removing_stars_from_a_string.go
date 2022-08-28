package contest

// Link: https://leetcode.com/problems/removing-stars-from-a-string
func removeStars(s string) string {
	runes := make([]rune, 0, len(s))
	for _, r := range s {
		if r == '*' {
			runes = runes[:len(runes)-1]
		} else {
			runes = append(runes, r)
		}
	}
	return string(runes)
}
