package projectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_totalFruits(t *testing.T) {
	assert.Equal(t, 3, totalFruit([]int{1, 2, 1}))
	assert.Equal(t, 3, totalFruit([]int{0, 1, 2, 2}))
	assert.Equal(t, 4, totalFruit([]int{1, 2, 3, 2, 2}))
	assert.Equal(t, 5, totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}
