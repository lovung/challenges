// Given an array of integers of size >= 2, find its two largest elements

// int64
// 3 elements [0, 1, 1, 1] -> [1, 1]
//

package coupang

import (
	"reflect"
	"testing"
)

func Test_findTwoLargest(t *testing.T) {
	type args struct {
		input []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			args: args{input: []int64{0, 1, 1, 1, 1}},
			want: []int64{1, 1},
		},
		{
			args: args{input: []int64{0, 1, 2, 3, 4, 5}},
			want: []int64{5, 4},
		},
		{
			args: args{input: []int64{6, 1, 2, 3, 4, 5}},
			want: []int64{6, 5},
		},
		{
			args: args{input: []int64{6, 1, 1, 1, 1, 1, 1}},
			want: []int64{6, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTwoLargest(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTwoLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
