package biweekly129

import "testing"

func Test_numberOfRightTriangles(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int64
	}{
		{
			grid: [][]int{
				{0, 0},
				{0, 1},
				{1, 1},
			},
			want: 1,
		},
		{
			grid: [][]int{
				{1, 0, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 1},
			},
			want: 9,
		},
		{
			grid: [][]int{
				{0, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			want: 2,
		},
		{
			grid: [][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfRightTriangles(tt.grid); got != tt.want {
				t.Errorf("numberOfRightTriangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
