package apr2024

import (
	"reflect"
	"testing"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			},
			want: []int{8, 12, 6, 10, 10, 10},
		},
		{
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}},
			},
			want: []int{15, 11, 9, 9, 11, 15},
		},
		{
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: []int{0},
		},
		{
			args: args{
				n:     2,
				edges: [][]int{{0, 1}},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDistancesInTree1(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfDistancesInTree() = %v, want %v", got, tt.want)
			}
			if got := sumOfDistancesInTree3(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfDistancesInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
