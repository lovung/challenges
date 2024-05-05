package contest394

func numberOfSpecialChars2(word string) int {
	index := make([]int, 256)
	for i := range index {
		index[i] = -1
	}
	for i := range word {
		if isLower(word[i]) || index[int(word[i])] < 0 {
			index[int(word[i])] = i
		}
	}
	const delta = 'a' - 'A'
	res := 0
	for i := 'A'; i <= 'Z'; i++ {
		if index[i] >= 0 && index[i+delta] >= 0 && index[i] > index[i+delta] {
			res++
		}
	}
	return res
}

func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}
