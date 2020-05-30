package maychallenge

func countBits(num int) []int {
	out := make([]int, num+1)
	count := 0
	current := 0
	for i := 0; i <= num; i++ {
		current = i
		for current > 0 {
			count += current & 1
			current >>= 1
		}
		out[i] = count
		count = 0
	}
	return out
}
