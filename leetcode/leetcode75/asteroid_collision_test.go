package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_asteroidCollision(t *testing.T) {
	assert.Equal(t, []int{5, 10}, asteroidCollision([]int{5, 10, -5}))
	assert.Equal(t, []int{}, asteroidCollision([]int{8, -8}))
	assert.Equal(t, []int{10}, asteroidCollision([]int{10, 2, -5}))
	assert.Equal(t, []int{-2, -2, -1}, asteroidCollision([]int{-2, 1, -2, -1}))
	assert.Equal(t, []int{-10}, asteroidCollision([]int{5, -10}))
	assert.Equal(t, []int{}, asteroidCollision([]int{10, 9, 5, -10}))
}
