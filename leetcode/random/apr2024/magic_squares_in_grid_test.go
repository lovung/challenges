package apr2024

import "testing"

func Test_numMagicSquaresInside(t *testing.T) {
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
				grid: [][]int{{7, 6, 2, 2, 4}, {4, 4, 9, 2, 10}, {9, 7, 8, 3, 10}, {8, 1, 9, 7, 5}, {7, 10, 4, 11, 6}},
			},
			want: 0,
		},
		{
			args: args{
				grid: [][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}},
			},
			want: 0,
		},
		{
			args: args{
				grid: [][]int{
					{4, 3, 8, 4},
					{9, 5, 1, 9},
					{2, 7, 6, 2}},
			},
			want: 1,
		},
		{
			args: args{
				grid: [][]int{
					{4, 3, 8, 4},
					{9, 6, 1, 9},
					{2, 7, 5, 2}},
			},
			want: 0,
		},
		{
			args: args{
				grid: [][]int{
					{4, 3, 8, 4},
					{9, 5, 1, 9},
					{6, 7, 2, 2}},
			},
			want: 0,
		},
		{
			args: args{
				grid: [][]int{
					{4, 3, 8, 4},
					{9, 5, 1, 9},
					{7, 2, 6, 2}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMagicSquaresInside(tt.args.grid); got != tt.want {
				t.Errorf("numMagicSquaresInside() = %v, want %v", got, tt.want)
			}
		})
	}
}
