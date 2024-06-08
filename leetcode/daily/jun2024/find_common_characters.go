package jun2024

func commonChars(words []string) []string {
	const a = 'a'
	cnt := [26]int{}
	for _, i := range words[0] {
		cnt[i-a]++
	}
	for _, w := range words {
		newCnt := [26]int{}
		for i := range w {
			newCnt[w[i]-a]++
		}
		for i := range 26 {
			cnt[i] = min(cnt[i], newCnt[i])
		}
	}
	res := make([]string, 0, 26)
	for i := range 26 {
		for range cnt[i] {
			res = append(res, string(byte(i)+'a'))
		}
	}
	return res
}
