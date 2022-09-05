package palindrome

// Link: https://leetcode.com/problems/palindrome-partitioning
func partition(s string) [][]string {
	curList := make([]string, 0)
	res := make([][]string, 0)
	dfsPartition(&res, curList, s)
	return res
}

func isPalindrome(s string) bool {
	n := len(s)
	n2 := n / 2
	for i := 0; i < n2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func dfsPartition(res *[][]string, curList []string, remainStr string) {
	if len(remainStr) == 0 {
		*res = append(*res, curList)
		return
	}
	for end := 1; end <= len(remainStr); end++ {
		if isPalindrome(remainStr[:end]) {
			cloneCurList := make([]string, len(curList), len(curList)+1)
			copy(cloneCurList, curList)
			cloneCurList = append(cloneCurList, remainStr[:end])
			dfsPartition(res, cloneCurList, remainStr[end:])
		}
	}
}
