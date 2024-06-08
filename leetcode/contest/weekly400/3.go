package weekly400

import "strings"

func clearStars(s string) string {
	const a = 'a'
	b := []byte(s)
	idx := [26][]int{}
	for i := range b {
		if b[i] == '*' {
			for j := range 26 {
				if len(idx[j]) > 0 {
					b[idx[j][len(idx[j])-1]] = '_'
					idx[j] = idx[j][:len(idx[j])-1]
					break
				}
			}
		} else {
			idx[b[i]-a] = append(idx[b[i]-a], i)
		}
	}
	sb := strings.Builder{}
	for i := range b {
		if b[i] == '*' || b[i] == '_' {
			continue
		}
		sb.WriteByte(b[i])
	}
	return sb.String()
}
