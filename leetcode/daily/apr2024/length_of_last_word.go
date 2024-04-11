package apr2024

// https://leetcode.com/problems/length-of-last-word
func lengthOfLastWord(s string) int {
	const space = ' '
	l, r := -1, -1
	for i := len(s) - 1; i >= 0; i-- {
		if r == -1 && s[i] != space {
			r = i
		} else if r != -1 && s[i] == space {
			l = i
			break
		}
	}
	return r - l
}
