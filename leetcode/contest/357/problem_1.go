package contest

// https://leetcode.com/problems/faulty-keyboard/
func finalString(s string) string {
	out := ""
	for i := range s {
		if s[i] == 'i' {
			out = reverseString(out)
			continue
		}
		out += string(s[i])
	}
	return out
}

func reverseString(s string) string {
	b := []byte(s)
	n2 := len(b) / 2
	n := len(b)
	for i := 0; i < n2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b)
}
