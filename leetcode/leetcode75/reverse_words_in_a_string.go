package leetcode75

import "strings"

// https://leetcode.com/problems/reverse-words-in-a-string
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	revereddWords := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) > 0 {
			revereddWords = append(revereddWords, words[i])
		}
	}
	return strings.Join(revereddWords, " ")
}
