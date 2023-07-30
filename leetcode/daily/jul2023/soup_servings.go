package jul2023

type pairRemain struct {
	a, b int // remain of each types
}

// https://leetcode.com/problems/soup-servings/
func soupServings(n int) float64 {
	if n%25 == 0 {
		n /= 25
	} else {
		n = n/25 + 1
	}
	cacheMap := make(map[pairRemain]float64)
	for i := 1; i <= n; i++ {
		if calculateProbability(i, i, cacheMap) > 1-1e-5 {
			return 1
		}
	}
	return calculateProbability(n, n, cacheMap)
}

func calculateProbability(ra, rb int, cache map[pairRemain]float64) float64 {
	switch {
	case ra <= 0 && rb <= 0:
		return 0.5
	case ra <= 0:
		return 1
	case rb <= 0:
		return 0
	}
	if val, ok := cache[pairRemain{ra, rb}]; ok {
		return val
	}

	prob := float64(0)
	prob += calculateProbability(ra-4, rb, cache)
	prob += calculateProbability(ra-3, rb-1, cache)
	prob += calculateProbability(ra-2, rb-2, cache)
	prob += calculateProbability(ra-1, rb-3, cache)
	prob /= 4

	cache[pairRemain{ra, rb}] = prob
	return prob
}
