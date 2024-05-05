package weekly396

func minAnagramLength(s string) int {
	b := []byte(s)
outter:
	for k := 1; k <= len(b); k++ {
		if len(b)%k != 0 {
			continue
		}
		firstK := count(b[:k])
		for i := k; i < len(b); i += k {
			nextK := count(b[i : i+k])
			if !compareCounter(firstK, nextK) {
				continue outter
			}
		}
		return k
	}
	return len(b)
}

func count(b []byte) []int {
	cnt := [26]int{}
	for i := range b {
		cnt[b[i]-'a']++
	}
	return cnt[:]
}

func compareCounter(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
