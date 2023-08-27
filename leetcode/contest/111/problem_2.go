package contest

// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
func canMakeSubsequence(str1 string, str2 string) bool {
	if len(str2) > len(str1) {
		return false
	}
	i, j := 0, 0
	for i < len(str2) && j < len(str1) {
		if str2[i] == str1[j] || str2[i] == increaseChar(str1[j]) {
			i++
			j++
			continue
		}
		j++
	}
	return i == len(str2)
}

func increaseChar(b byte) byte {
	if b == 'z' {
		return 'a'
	}
	return b + 1
}
