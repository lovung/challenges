package grind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}
