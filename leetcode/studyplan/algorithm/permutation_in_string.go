package algorithm

func checkInclusion(s1 string, s2 string) bool {
	n := int('z' - 'a')
	chars := make([]int, n)
	for i := range s1 {
		chars[s1[i]-'a']++
	}
	for i := range s2 {
		chars[s2[i]-'a']--
	}
	for i := range chars {
		if chars[i] > 0 {
			return false
		}
	}
	return true
}
