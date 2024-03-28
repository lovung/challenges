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
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{2, 2, 2, 2, 2}, 2},
		{[]int{1, 2, 2, 2, 2}, 2},
		{[]int{1, 1}, 1},
	}
	for _, c := range testCase {
		t.Run("sol_1_hashmap", func(t *testing.T) {
			input := make([]int, len(c.input))
			copy(input, c.input)
			assert.Equal(t, c.expect, findDuplicate_hashmap(input))
		})
		t.Run("sol_2_bubble_sort", func(t *testing.T) {
			input := make([]int, len(c.input))
			copy(input, c.input)
			assert.Equal(t, c.expect, findDuplicate_bubbleSort(input))
		})
		t.Run("sol_3_tortoiseAndHare", func(t *testing.T) {
			input := make([]int, len(c.input))
			copy(input, c.input)
			assert.Equal(t, c.expect, findDuplicate_tortoiseAndHare(input))
		})

	}
}
