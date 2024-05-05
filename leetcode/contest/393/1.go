package contest393

import "strconv"

func findLatestTime(s string) string {
	const first, second, third, forth = 0, 1, 3, 4
	max1, max2, max3, max4 := 1, 9, 5, 9
	const q = '?'
	res := ""
	switch s[first] {
	case q:
		if s[second] == q || s[second] <= '1' {
			res += strconv.Itoa(max1)
			max2 = 1
		} else {
			res += "0"
		}
	case '1':
		max2 = 1
		fallthrough
	default:
		res += string(s[first])
	}
	switch s[second] {
	case q:
		res += strconv.Itoa(max2)
	default:
		res += string(s[second])
	}
	res += ":"
	switch s[third] {
	case q:
		res += strconv.Itoa(max3)
	default:
		res += string(s[third])
	}
	switch s[forth] {
	case q:
		res += strconv.Itoa(max4)
	default:
		res += string(s[forth])
	}
	return res
}
