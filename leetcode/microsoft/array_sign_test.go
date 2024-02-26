package microsoft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arraySign(t *testing.T) {
	input := [][]int{
		{-1, -2, -3, -4, 3, 2, 1},
		{1, 5, 0, 2, -3},
		{-1, 1, -1, 1, -1},
		{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24},
	}
	expected := []int{1, 0, -1, -1}
	for i, c := range input {
		assert.Equal(t, expected[i], arraySign(c))
	}
}
