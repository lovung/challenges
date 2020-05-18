package maychallenge

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3333/
 */

func checkInclusion(s1 string, s2 string) bool {
	const a = 'a'
	lS1 := len(s1)
	lS2 := len(s2)
	s1Arr := make([]int, 26)
	s2Arr := make([]int, 26)

	for i := 0; i < lS1; i++ {
		s1Arr[s1[i]-a]++
	}
	for i := 0; i < lS2; i++ {
		s2Arr[s2[i]-a]++
		if i >= lS1 {
			s2Arr[s2[i-lS1]-a]--
		}
		if verify2CharArray(s2Arr, s1Arr) {
			return true
		}
	}
	return false
}

func verify2CharArray(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
