package jul2023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredictTheWinner(t *testing.T) {
	assert.False(t, PredictTheWinner([]int{1, 5, 2}))
	assert.True(t, PredictTheWinner([]int{1, 5, 233, 7}))
	assert.False(t, PredictTheWinner2([]int{1, 5, 2}))
	assert.True(t, PredictTheWinner2([]int{1, 5, 233, 7}))
	assert.False(t, PredictTheWinner3([]int{1, 5, 2}))
	assert.True(t, PredictTheWinner3([]int{1, 5, 233, 7}))
	assert.False(t, PredictTheWinner4([]int{1, 5, 2}))
	assert.True(t, PredictTheWinner4([]int{1, 5, 233, 7}))
}
