package litmus

func findMaxSum(numbers []int) int {
	firstMax, secondMax := 0, 0
	for i := range numbers {
		if firstMax <= numbers[i] {
			secondMax = firstMax
			firstMax = numbers[i]
		} else if secondMax <= numbers[i] {
			secondMax = numbers[i]
		}
	}
	return firstMax + secondMax
}
