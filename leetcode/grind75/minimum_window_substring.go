package grind75

// Link: https://leetcode.com/problems/minimum-window-substring/
const k = 'z' - 'A' + 1
const A = 'A'

func minWindow(s string, t string) string {
	// BruceForce solution with Counter:
	// 	k = 'z'-'A'
	// 	count the [k]counter for t
	// 	because we want to find minLen substring					-> O(N)
	// 	check all substrings of s from the len == 1 increasing  	-> O(M^2)
	//		each substring can calculate the [k]counter 			-> O(1) using counter
	//		compare if all item in counterS >= same index in counterT
	// total: O(N + M^2)

	cntT := countChar(t)

	// l is len of substring
	cntS := make([]int, k)
	for l := 1; l <= len(s); l++ {
		_cntS := make([]int, k)
		cntS[s[l-1]-A]++
		copy(_cntS, cntS)
		// i is the start index of substring
		for i := 0; i+l <= len(s); i++ {
			if checkCanCover(_cntS, cntT) {
				return s[i : i+l]
			}
			if i+l < len(s) {
				_cntS[s[i]-A]--
				_cntS[s[i+l]-A]++
			}
		}
	}
	return ""
}

func countChar(t string) []int {
	cntT := make([]int, k)
	for i := range t {
		cntT[t[i]-A]++
	}
	return cntT
}

func checkCanCover(cntS, cntT []int) bool {
	for i := range cntS {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

// BigO: O(M+N)
func minWindow2(s string, t string) string {
	// Sliding window
	cntT := countChar(t)

	// l is len of substring
	cntS := make([]int, k)
	l, r := 0, 0
	cntS[s[r]-A]++
	minWindowL := 0
	minWindowR := len(s)

	for r < len(s) {
		if checkCanCover(cntS, cntT) {
			if l == r {
				return s[l : r+1]
			}
			if minWindowR-minWindowL+1 > r-l+1 {
				minWindowL = l
				minWindowR = r
			}
			cntS[s[l]-A]--
			l++
		} else {
			r++
			if r < len(s) {
				cntS[s[r]-A]++
			}
		}
	}
	if minWindowR-minWindowL+1 > len(s) {
		return ""
	}
	return s[minWindowL : minWindowR+1]
}
