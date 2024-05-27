package weekly399

import (
	"strconv"
	"strings"
)

// https://leetcode.com/contest/weekly-contest-399/problems/string-compression-iii/description/
func compressedString(word string) string {
	comp := strings.Builder{}
	for i := 0; i < len(word); {
		j := i + 1
		for ; j < i+9 && j < len(word); j++ {
			if word[j] != word[i] {
				break
			}
		}
		comp.WriteString(strconv.Itoa(j - i))
		comp.WriteByte(word[i])
		i = j
	}
	return comp.String()
}
