package contest

import "github.com/lovung/ds/maths"

// https://leetcode.com/problems/number-of-beautiful-integers-in-the-range/
func numberOfBeautifulIntegers(low int, high int, k int) int {
	cnt := 0
	i := (low / k) * k
	if low%k != 0 {
		i += k
	}
	for i <= high {
		if isBeatiful(i) {
			cnt++
		}
		if n := getDigitCnt(i); n%2 != 0 {
			jump := maths.Max((powTen(n)-i)/k, 1)
			i += k * jump
		} else {
			i += k
		}
	}
	return cnt
}

func powTen(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}

func getDigitCnt(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		n /= 10
	}
	return cnt
}

func isBeatiful(n int) bool {
	oddCnt := 0
	evenCnt := 0
	for n > 0 {
		if n%2 == 0 {
			evenCnt++
		} else {
			oddCnt++
		}
		n /= 10
	}
	return evenCnt == oddCnt
}
