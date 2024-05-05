package apr2024

import (
	"reflect"
	"testing"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		n         int
		edges     [][]int
		disappear []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:         3,
				edges:     [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}},
				disappear: []int{1, 3, 5},
			},
			want: []int{0, 2, 3},
		},
		{
			args: args{
				n:         3,
				edges:     [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}},
				disappear: []int{1, 1, 5},
			},
			want: []int{0, -1, 4},
		},
		{
			args: args{
				n: 3,
				edges: [][]int{
					{2, 0, 9}, {1, 0, 5},
					{2, 2, 4}, {0, 1, 10},
					{1, 1, 10}, {1, 1, 10},
					{2, 2, 10}, {1, 1, 10}},
				disappear: []int{4, 13, 19},
			},
			want: []int{0, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.args.n, tt.args.edges, tt.args.disappear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
