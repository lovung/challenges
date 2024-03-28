package mar2024

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_reorderList(t *testing.T) {
	testcase := []struct {
		input, want []int
	}{
		{
			input: []int{1, 2, 3, 4},
			want:  []int{1, 4, 2, 3},
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 5, 2, 4, 3},
		},
		{
			input: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			input: []int{1},
			want:  []int{1},
		},
		{
			input: []int{},
			want:  []int{},
		},
	}
	for _, tc := range testcase {
		head := lists.NewLinkedListFromValues(tc.input)
		reorderList(head)
		assert.Equal(t,
			lists.NewLinkedListFromValues(tc.want),
			head,
		)
	}
}
