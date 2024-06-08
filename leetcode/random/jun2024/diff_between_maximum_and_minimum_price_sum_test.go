package jun2024

import "testing"

func Test_maxOutput(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		price []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}},
				price: []int{9, 8, 7, 6, 10, 5},
			},
			want: 24,
		},
		{
			args: args{
				n:     8,
				edges: [][]int{{1, 7}, {2, 3}, {4, 0}, {5, 7}, {6, 3}, {3, 0}, {0, 7}},
				price: []int{4, 5, 6, 2, 2, 7, 7, 8},
			},
			want: 21,
		},
		{
			args: args{
				n:     3,
				edges: [][]int{{1, 0}, {2, 0}},
				price: []int{3, 2, 2},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOutput(tt.args.n, tt.args.edges, tt.args.price); got != tt.want {
				t.Errorf("maxOutput() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOutput1(tt.args.n, tt.args.edges, tt.args.price); got != tt.want {
				t.Errorf("maxOutput1() = %v, want %v", got, tt.want)
			}
		})
	}
}
