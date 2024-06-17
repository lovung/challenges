package weekly402

import (
	"reflect"
	"testing"
)

func Test_countOfPeaks(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			args: args{
				nums:    []int{5, 10, 7, 4, 9, 7, 3, 8},
				queries: [][]int{{1, 3, 6}, {1, 5, 7}, {1, 0, 4}},
			},
			want: []int{1, 0, 1},
		},
		{
			args: args{
				nums:    []int{4, 9, 4, 10, 7},
				queries: [][]int{{2, 3, 2}, {2, 1, 3}, {1, 2, 3}},
			},
			want: []int{0},
		},
		{
			args: args{
				nums:    []int{3, 1, 4, 2, 5},
				queries: [][]int{{2, 3, 4}, {1, 0, 4}},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfPeaks(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countOfPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
