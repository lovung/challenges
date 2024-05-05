package weekly396

func isValid(word string) bool {
	var (
		vowels = map[byte]bool{
			'a': true,
			'e': true,
			'i': true,
			'o': true,
			'u': true,
		}
	)
	if len(word) < 3 {
		return false
	}
	vowelCnt, consonantCnt := 0, 0
	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			continue
		}
		if 'a' <= word[i] && word[i] <= 'z' {
			if vowels[word[i]] {
				vowelCnt++
			} else {
				consonantCnt++
			}
		} else if 'A' <= word[i] && word[i] <= 'Z' {
			if vowels[word[i]-'A'+'a'] {
				vowelCnt++
			} else {
				consonantCnt++
			}
		} else {
			return false
		}
	}
	return vowelCnt > 0 && consonantCnt > 0
}
