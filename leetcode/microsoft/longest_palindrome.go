package microsoft

func longestPalindrome(s string) string {
	b := []byte(s)
	var ret []byte
	for i := range b {
		j := checkPalindromicOdd(b, i)
		if j+j+1 > len(ret) {
			ret = b[i-j : i+j+1]
		}
		j = checkPalindromicEven(b, i)
		if j+j > len(ret) {
			ret = b[i-j+1 : i+j+1]
		}
	}
	return string(ret)
}

func checkPalindromicOdd(b []byte, i int) int {
	j := 0
	for ; i-j >= 0 && i+j < len(b); j++ {
		if b[i+j] != b[i-j] {
			return j - 1
		}
	}
	return j - 1
}

func checkPalindromicEven(b []byte, i int) int {
	j := 0
	for ; i-j >= 0 && i+j+1 < len(b); j++ {
		if b[i+1+j] != b[i-j] {
			return j
		}
	}
	return j
}
