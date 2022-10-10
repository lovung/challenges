package oct2022

// Link: https://leetcode.com/problems/break-a-palindrome/
func breakPalindrome(palindrome string) string {
	midIdx := -1
	if len(palindrome)%2 != 0 {
		midIdx = len(palindrome) / 2
	}
	// replace the first not 'a' and not in middle to 'b'
	// abccba -> aaccba
	for i := range palindrome {
		if palindrome[i] != 'a' && i != midIdx {
			return changeCharAtITo(palindrome, i, 'a')
		}
	}

	// if can't find, replace last 'a' not in middle to 'b
	// aba -> abb
	// aaaaa -> aaaab
	for i := len(palindrome) - 1; i >= 0; i-- {
		if palindrome[i] == 'a' && i != midIdx {
			return changeCharAtITo(palindrome, i, 'b')
		}
	}

	return ""
}

func changeCharAtITo(s string, at int, to byte) string {
	s2 := []byte(s)
	s2[at] = to
	return string(s2)
}
