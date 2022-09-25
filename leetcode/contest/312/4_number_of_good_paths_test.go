package contest

import "testing"

func Test_numberOfGoodPaths(t *testing.T) {
	type args struct {
		vals  []int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				vals:  []int{1, 3, 2, 1, 3},
				edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
			},
			want: 6,
		},
		{
			args: args{
				vals:  []int{1, 1, 2, 2, 3},
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}},
			},
			want: 7,
		},
		{
			args: args{
				vals:  []int{1},
				edges: [][]int{},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfGoodPaths(tt.args.vals, tt.args.edges); got != tt.want {
				t.Errorf("numberOfGoodPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
