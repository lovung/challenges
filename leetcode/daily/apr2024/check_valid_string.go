package apr2024

// https://leetcode.com/problems/valid-parenthesis-string/description
func checkValidString(s string) bool {
	minVal, maxVal := 0, 0
	for i := range s {
		switch s[i] {
		case '(':
			minVal++
			maxVal++
		case ')':
			minVal = max(0, minVal-1)
			maxVal--
			if maxVal < 0 {
				return false
			}
		default: // '*'
			minVal = max(0, minVal-1)
			maxVal++
		}
	}
	return minVal <= 0 && maxVal >= 0
}

func checkValidString1(s string) bool {
	bs := []byte(s)
	return checkValid(bs, 0)
}

func checkValid(bs []byte, cnt int) bool {
	if len(bs) == 0 {
		return cnt == 0
	}
	if cnt < 0 {
		return false
	}

	switch bs[0] {
	case '(':
		return checkValid(bs[1:], cnt+1)
	case ')':
		return checkValid(bs[1:], cnt-1)
	default: // '*'
		return checkValid(bs[1:], cnt+1) ||
			checkValid(bs[1:], cnt) ||
			checkValid(bs[1:], cnt-1)
	}
}
