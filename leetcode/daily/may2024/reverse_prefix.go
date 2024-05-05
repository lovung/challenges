package may2024

import "strings"

func reversePrefix(word string, ch byte) string {
	idx := strings.IndexByte(word, ch)
	return reverse(word[:idx+1]) + word[idx+1:]
}

func reversePrefix2(word string, ch byte) string {
	k := len(word) - 1
	prefixReversed := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		prefixReversed[k] = word[i]
		if word[i] == ch {
			return string(prefixReversed[k:]) + word[i+1:]
		}
		k--
	}
	return word
}

func reverse(s string) string {
	sb := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
