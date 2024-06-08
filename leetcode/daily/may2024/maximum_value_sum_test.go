package may2024

import "testing"

func Test_maximumValueSum(t *testing.T) {
	type args struct {
		nums  []int
		k     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums:  []int{1, 2, 1},
				k:     3,
				edges: [][]int{{0, 1}, {0, 2}},
			},
			want: 6,
		},
		{
			args: args{
				nums:  []int{2, 3},
				k:     7,
				edges: [][]int{{0, 1}},
			},
			want: 9,
		},
		{
			args: args{
				nums:  []int{7, 7, 7, 7, 7},
				k:     15,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumValueSum(tt.args.nums, tt.args.k, tt.args.edges); got != tt.want {
				t.Errorf("maximumValueSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
