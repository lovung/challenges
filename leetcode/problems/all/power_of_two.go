package problems

func isPowerOfTwo(n int) bool {
	countBit := 0
	if n < 0 {
		return false
	}
	for n != 0 {
		if n&1 == 1 {
			countBit++
		}
		n >>= 1
	}
	return countBit == 1
}
