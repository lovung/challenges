package coupang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CutInitialPortion(t *testing.T) {
	tuples := []Tuple{
		{2, 3}, {5, 4},
	}
	assert.Equal(t, []Tuple{}, CutInitialPortion([]Tuple{}, 0))
	assert.Equal(t, []Tuple{{2, 3}, {5, 4}}, CutInitialPortion(tuples, 0))
	assert.Equal(t, []Tuple{{1, 2}, {1, 3}, {5, 4}}, CutInitialPortion(tuples, 1))
	assert.Equal(t, []Tuple{{1, 1}, {1, 3}, {5, 4}}, CutInitialPortion(tuples, 2))
	assert.Equal(t, []Tuple{{1, 3}, {5, 4}}, CutInitialPortion(tuples, 3))
	assert.Equal(t, []Tuple{{1, 2}, {5, 4}}, CutInitialPortion(tuples, 4))
	assert.Equal(t, []Tuple{{1, 1}, {5, 4}}, CutInitialPortion(tuples, 5))
	assert.Equal(t, []Tuple{{5, 4}}, CutInitialPortion(tuples, 6))
}
