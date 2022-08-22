package easy

// Link: https://leetcode.com/problems/to-lower-case/
func toLowerCase(s string) string {
	var result string
	const delta byte = byte('a') - byte('A')
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result += string(s[i] + delta)
		} else {
			result += string(s[i])
		}
	}
	return result
}
