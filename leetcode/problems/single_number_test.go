package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
	assert.Equal(t, 4, singleNumber([]int{4, 1, 2, 1, 2}))
	assert.Equal(t, 1, singleNumber([]int{1}))
	assert.Equal(t, 0, singleNumber([]int{}))
}

func Test_singleNumber2(t *testing.T) {
	assert.Equal(t, 3, singleNumber2([]int{2, 2, 3, 2}))
	assert.Equal(t, 99, singleNumber2([]int{0, 1, 0, 1, 0, 1, 99}))
	assert.Equal(t, 0, singleNumber2([]int{}))
}

func Test_singleNumber3(t *testing.T) {
	assert.ElementsMatch(t, []int{3, 5}, singleNumber3([]int{1, 2, 1, 3, 2, 5}))
	assert.ElementsMatch(t, []int{-1, 0}, singleNumber3([]int{-1, 0}))
	assert.ElementsMatch(t, []int{}, singleNumber3([]int{}))
}
