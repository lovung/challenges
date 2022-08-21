package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[5,4,7,5,3,2]",
			args: args{
				nums: []int{5, 4, 7, 5, 3, 2},
			},
			want: []int{5, 5, 2, 3, 4, 7},
		},
		{
			name: "[1, 2, 3]",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "[1, 3, 2]",
			args: args{
				nums: []int{1, 3, 2},
			},
			want: []int{2, 1, 3},
		},
		{
			name: "[2, 3, 1]",
			args: args{
				nums: []int{2, 3, 1},
			},
			want: []int{3, 1, 2},
		},
		{
			name: "[3, 2, 1]",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "[1, 1, 5]",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: []int{1, 5, 1},
		},
		{
			name: "[1, 5, 1]",
			args: args{
				nums: []int{1, 5, 1},
			},
			want: []int{5, 1, 1},
		},
		{
			name: "[1, 9, 8, 12, 3]",
			args: args{
				nums: []int{1, 9, 8, 12, 3},
			},
			want: []int{1, 9, 12, 3, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
