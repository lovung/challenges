package weekly402

func countCompleteDayPairs(hours []int) int {
	mod := [24]int{}
	res := 0
	for _, h := range hours {
		modH := h % 24
		if modH == 0 {
			res += mod[0]
		} else {
			res += mod[24-modH]
		}
		mod[modH]++
	}
	return res
}
