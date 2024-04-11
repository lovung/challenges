package apr2024

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
func minRemoveToMakeValid(s string) string {
	level, res := 0, make([]byte, 0, len(s))
	for i := range s {
		if level == 0 && s[i] == ')' {
			continue
		}
		res = append(res, s[i])
		if s[i] == '(' {
			level++
		} else if s[i] == ')' {
			level--
		}
	}
	for j := len(res) - 1; j >= 0 && level > 0; j-- {
		if res[j] == '(' {
			copy(res[j:], res[j+1:])
			res = res[:len(res)-1]
			level--
		}
	}
	return string(res)
}
