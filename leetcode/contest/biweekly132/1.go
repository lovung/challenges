package biweekly132

// https://leetcode.com/problems/clear-digits/description/
func clearDigits(s string) string {
	res := make([]byte, 0, len(s))
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' && len(res) > 0 {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
