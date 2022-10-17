package oct2022

// Link: https://leetcode.com/problems/check-if-the-sentence-is-pangram/submissions/
func checkIfPangram(sentence string) bool {
	var cnt [26]int
	for i := range sentence {
		cnt[sentence[i]-'a']++
	}

	for i := range cnt {
		if cnt[i] == 0 {
			return false
		}
	}
	return true
}
