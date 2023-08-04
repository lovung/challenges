package leetcode75

// https://leetcode.com/problems/asteroid-collision
func asteroidCollision(asteroids []int) []int {
	for i := range asteroids {
		runByDirection(asteroids, i)
	}
	return removeZero(asteroids)
}

func removeZero(slice []int) []int {
	ret := make([]int, 0, len(slice))
	for i := range slice {
		if slice[i] != 0 {
			ret = append(ret, slice[i])
		}
	}
	return ret
}

func runByDirection(asteroids []int, start int) {
	size := asteroids[start]
	if size > 0 {
		for i := start + 1; i < len(asteroids); i++ {
			switch {
			case asteroids[i] > 0:
				return
			case asteroids[i] < 0:
				if -asteroids[i] > size {
					asteroids[start] = 0
					return
				} else if -asteroids[i] < size {
					asteroids[i] = 0
				} else {
					asteroids[start], asteroids[i] = 0, 0
					return
				}
			}
		}
	} else if size < 0 {
		for i := start - 1; i >= 0; i-- {
			switch {
			case asteroids[i] < 0:
				return
			case asteroids[i] > 0:
				if asteroids[i] > -size {
					asteroids[start] = 0
					return
				} else if asteroids[i] < -size {
					asteroids[i] = 0
				} else {
					asteroids[start], asteroids[i] = 0, 0
					return
				}
			}
		}
	}
}
