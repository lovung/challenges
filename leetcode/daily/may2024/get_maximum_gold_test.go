package may2024

import "testing"

func Test_getMaximumGold(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]int{{0, 0, 19, 5, 8}, {11, 20, 14, 1, 0}, {0, 0, 1, 1, 1}, {0, 2, 0, 2, 0}},
			},
			want: 77,
		},
		{
			args: args{
				grid: [][]int{{0, 0, 1}, {1, 1, 0}, {1, 1, 0}},
			},
			want: 4,
		},
		{
			args: args{
				grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}},
			},
			want: 24,
		},
		{
			args: args{
				grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}},
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGold(tt.args.grid); got != tt.want {
				t.Errorf("getMaximumGold() = %v, want %v", got, tt.want)
			}
		})
	}
}
