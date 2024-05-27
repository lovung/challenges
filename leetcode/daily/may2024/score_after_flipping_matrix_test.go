package may2024

import "testing"

func Test_matrixScore(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}},
			want: 39,
		},
		{
			grid: [][]int{{1, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}},
			want: 38,
		},
		{
			grid: [][]int{{0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixScore2(tt.grid); got != tt.want {
				t.Errorf("matrixScore2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixScore(tt.grid); got != tt.want {
				t.Errorf("matrixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
