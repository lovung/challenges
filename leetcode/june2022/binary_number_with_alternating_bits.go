package june2022

func hasAlternatingBits(n int) bool {
	bitMark := 1
	isLastBitZero := (n & bitMark) == 0
	for n != 0 {
		n >>= 1
		isCurrentBitZero := (n & bitMark) == 0
		if isCurrentBitZero == isLastBitZero {
			return false
		}
		isLastBitZero = isCurrentBitZero
	}
	return true
}
