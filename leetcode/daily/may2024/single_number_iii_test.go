package may2024

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				[]int{1, 2, 1, 3, 2, 5},
			},
			want: []int{3, 5},
		},
		{
			args: args{
				[]int{1, 2, 1, 3, 2, 7, 7, 5, 5, 21, 11, 3},
			},
			want: []int{11, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.args.nums)
			sort.Ints(got)
			assert.EqualValues(t, tt.want, got)
		})
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber2(tt.args.nums)
			sort.Ints(got)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
