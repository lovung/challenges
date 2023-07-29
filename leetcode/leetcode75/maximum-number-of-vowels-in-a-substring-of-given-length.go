package leetcode75

func maxVowels(s string, k int) int {
	cnt := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			cnt++
		}
	}
	maxVowel := cnt
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) && !isVowel(s[i-k]) {
			cnt++
		} else if !isVowel(s[i]) && isVowel(s[i-k]) {
			cnt--
		}
		if maxVowel < cnt {
			maxVowel = cnt
		}
	}
	return maxVowel
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
