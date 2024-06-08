package weekly400

func minimumChairs(s string) int {
	res := 0
	cur := 0
	for i := range s {
		if s[i] == 'E' {
			cur++
		} else {
			cur--
		}
		res = max(res, cur)
	}
	return res
}
