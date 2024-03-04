package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicate(t *testing.T) {
	type test struct {
		input  []int
		expect int
	}
	testCase := []test{
		{[]int{}, 0},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{2, 2, 2, 2, 2}, 2},
		{[]int{1, 2, 2, 2, 2}, 2},
	}
	for _, c := range testCase {
		input := make([]int, len(c.input))
		copy(input, c.input)
		assert.Equal(t, c.expect, findDuplicate(input))
		input2 := make([]int, len(c.input))
		copy(input2, c.input)
		assert.Equal(t, c.expect, findDuplicate1(input2))
	}
}
