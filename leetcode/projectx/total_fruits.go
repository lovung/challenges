package projectx

// https://leetcode.com/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {
	const baskets = 2

	typesCounter := make(map[int]int)
	numOfTypesInRange := 0
	maxSubArrLen := 0

	for l, r := 0, 0; r < len(fruits); r++ {
		for ; numOfTypesInRange > baskets; l++ {
			typesCounter[fruits[l]]--
			if typesCounter[fruits[l]] == 0 {
				numOfTypesInRange--
			}
		}
		typesCounter[fruits[r]]++
		if typesCounter[fruits[r]] == 1 {
			numOfTypesInRange++
		}
		if numOfTypesInRange <= baskets {
			maxSubArrLen = max(maxSubArrLen, r-l+1)
		}
	}
	return maxSubArrLen
}
