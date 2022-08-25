package easy

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Link: https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/
func Test_check(t *testing.T) {
	type cases struct {
		in   []int
		want bool
	}
	c := []cases{
		{
			in:   []int{3, 4, 5, 1, 2},
			want: true,
		},
		{
			in:   []int{2, 1, 3, 4},
			want: false,
		},
		{
			in:   []int{2, 1, 3, 1, 4},
			want: false,
		},
		{
			in:   []int{1, 2, 3},
			want: true,
		},
		{
			in:   []int{6, 10, 6},
			want: true,
		},
	}
	for i := range c {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, c[i].want, check(c[i].in))
		})
	}
}
