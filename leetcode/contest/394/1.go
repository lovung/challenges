package contest394

func numberOfSpecialChars1(word string) int {
	cnt := make([]int, 256)
	for i := range word {
		cnt[int(word[i])]++
	}
	const delta = 'a' - 'A'
	res := 0
	for i := 'A'; i <= 'Z'; i++ {
		if cnt[i] > 0 && cnt[i+delta] > 0 {
			res++
		}
	}
	return res
}
