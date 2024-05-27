package weekly398

import (
	"reflect"
	"testing"
)

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{
				nums:    []int{1, 1},
				queries: [][]int{{0, 1}},
			},
			want: []bool{false},
		},
		{
			args: args{
				nums:    []int{4, 1, 4, 5, 1, 2, 8, 2, 8, 1, 2, 8},
				queries: [][]int{{2, 6}},
			},
			want: []bool{false},
		},
		{
			args: args{
				nums:    []int{4, 3, 1, 6},
				queries: [][]int{{0, 2}, {2, 3}},
			},
			want: []bool{false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
