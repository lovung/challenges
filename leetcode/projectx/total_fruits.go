package projectx

// https://leetcode.com/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {
	const baskets = 2

	typesCounter := make(map[int]int)
	numOfTypes := 0
	maxFruitTypes := 0

	for l, r := 0, 0; r < len(fruits); r++ {
		for ; numOfTypes > baskets; l++ {
			typesCounter[fruits[l]]--
			if typesCounter[fruits[l]] == 0 {
				numOfTypes--
			}
		}
		typesCounter[fruits[r]]++
		if typesCounter[fruits[r]] == 1 {
			numOfTypes++
		}
		if numOfTypes <= baskets {
			maxFruitTypes = max(maxFruitTypes, r-l+1)
		}
	}
	return maxFruitTypes
}
