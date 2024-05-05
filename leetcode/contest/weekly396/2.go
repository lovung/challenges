package weekly396

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	parts := divides(word, k)
	cnt := make(map[string]int)
	maxCnt := 0
	for _, p := range parts {
		cnt[p]++
		maxCnt = max(maxCnt, cnt[p])
	}
	return len(parts) - maxCnt
}

func divides(word string, k int) []string {
	res := make([]string, 0, len(word)/k)
	for i := 0; i < len(word); i += k {
		res = append(res, word[i:i+k])
	}
	return res
}
