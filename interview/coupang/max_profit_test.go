package coupang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert.EqualValues(t, 9, maximumProfit([]int32{5, 3, 4, 6, 2, 5}))
	assert.EqualValues(t, 6, maximumProfit([]int32{5, 3, 4, 6, 2}))
	assert.EqualValues(t, 0, maximumProfit([]int32{7, 7, 7, 7, 7, 7}))
}
