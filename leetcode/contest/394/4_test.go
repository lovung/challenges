package contest394

import (
	"reflect"
	"testing"
)

func Test_findAnswer(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{

				n:     6,
				edges: [][]int{{0, 1, 4}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3}, {1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2}},
			},
			want: []bool{true, true, true, false, true, true, true, false},
		},
		{
			args: args{

				n:     6,
				edges: [][]int{{0, 1, 5}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3}, {1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2}},
			},
			want: []bool{false, true, true, false, true, true, true, false},
		},
		{
			args: args{
				n:     4,
				edges: [][]int{{2, 0, 1}, {0, 1, 1}, {0, 3, 4}, {3, 2, 2}},
			},
			want: []bool{true, false, false, true},
		},
		{
			args: args{
				n: 7,
				edges: [][]int{
					{5, 1, 10}, {2, 5, 3},
					{6, 4, 4}, {0, 5, 10}, {4, 1, 2}},
			},
			want: []bool{true, false, true, true, true},
		},
		{
			args: args{
				n:     6,
				edges: [][]int{{0, 1, 4}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3}, {1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2}},
			},
			want: []bool{true, true, true, false, true, true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnswer(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}
