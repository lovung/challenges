package contest392

import (
	"reflect"
	"testing"
)

func Test_minimumCost(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		query [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n: 10,
				edges: [][]int{
					{0, 1, 0}, {6, 8, 4}, {8, 1, 7},
				},
				query: [][]int{{0, 1}, {0, 9}, {2, 7}, {9, 4}, {3, 5}, {2, 6}},
			},
			want: []int{0, -1, -1, -1, -1, -1},
		},
		{
			args: args{
				n: 3,
				edges: [][]int{
					{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1},
				},
				query: [][]int{{1, 2}},
			},
			want: []int{0},
		},
		{
			args: args{
				n: 8,
				edges: [][]int{
					{0, 4, 7}, {3, 5, 1}, {1, 3, 5}, {1, 5, 1},
				},
				query: [][]int{{0, 4}, {1, 5}, {3, 0}, {3, 3}, {3, 2}, {2, 0}, {7, 7}, {7, 0}},
			},
			want: []int{7, 1, -1, 0, -1, -1, 0, -1},
		},
		{
			args: args{
				n: 8,
				edges: [][]int{
					{0, 1, 7}, {1, 3, 7}, {1, 2, 1},
				},
				query: [][]int{{0, 3}, {3, 4}},
			},
			want: []int{1, -1},
		},
		{
			args: args{
				n:     9,
				edges: [][]int{{5, 4, 9}, {6, 1, 1}, {6, 5, 7}, {5, 7, 4}, {5, 7, 8}, {8, 3, 0}, {4, 0, 8}, {3, 7, 0}, {1, 5, 3}},
				query: [][]int{{3, 0}},
			},
			want: []int{0},
		},
		{
			args: args{
				n:     8,
				edges: [][]int{{0, 3, 5}, {0, 3, 2}, {4, 5, 7}, {3, 1, 7}, {5, 2, 5}, {1, 0, 0}},
				query: [][]int{{4, 5}},
			},
			want: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.n, tt.args.edges, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
