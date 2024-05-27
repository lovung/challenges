package may2024

func partition(s string) [][]string {
	res := make([][]string, 0)
	partitionRecursive(s, &res, nil)
	return res
}

func partitionRecursive(s string, res *[][]string, prevPalindrome []string) {
	if len(s) == 0 {
		*res = append(*res, prevPalindrome)
		return
	}
	for i := 1; i <= len(s); i++ {
		curS := s[:i]
		if isPalindrome(curS) {
			newPalindrome := make([]string, len(prevPalindrome)+1)
			copy(newPalindrome, prevPalindrome)
			newPalindrome[len(newPalindrome)-1] = s[:i]
			partitionRecursive(s[i:], res, newPalindrome)
		}
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
