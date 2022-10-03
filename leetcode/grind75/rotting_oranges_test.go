package grind75

import "testing"

func Test_orangesRotting(t *testing.T) {
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
				grid: [][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				},
			},
			want: 4,
		},
		{
			args: args{
				grid: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{0, 1, 2},
				},
			},
			want: 4,
		},
		{
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				},
			},
			want: -1,
		},
		{
			args: args{
				grid: [][]int{
					{0, 2},
				},
			},
			want: 0,
		},
		{
			args: args{
				grid: [][]int{
					{1},
					{1},
					{1},
					{1},
				},
			},
			want: -1,
		},
		{
			args: args{
				grid: [][]int{
					{0},
				},
			},
			want: 0,
		},
		{
			args: args{
				grid: [][]int{
					{2},
					{2},
					{1},
					{0},
					{1},
					{1},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
