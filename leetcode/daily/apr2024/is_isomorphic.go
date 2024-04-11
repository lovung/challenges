package apr2024

// https://leetcode.com/problems/isomorphic-strings/description/
func isIsomorphic(s string, t string) bool {
	return isUniIsomorphic(s, t) && isUniIsomorphic(t, s)
}

func isUniIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	for i := range s {
		if e, ok := m[s[i]]; ok && e != t[i] {
			return false
		} else {
			m[s[i]] = t[i]
		}
	}
	return true
}
