package grind75

import "github.com/lovung/ds/maths"

func maxArea(height []int) int {
	length := len(height)
	square := 0
	for i := 0; i < length; i++ {
		if height[i] == 0 {
			continue
		}
		if (length-i)*height[i] < square {
			continue
		}
		for j := square / height[i]; j < length; j++ {
			square = maths.Max(square, (j-i)*maths.Min(height[i], height[j]))
		}
	}
	return square
}
