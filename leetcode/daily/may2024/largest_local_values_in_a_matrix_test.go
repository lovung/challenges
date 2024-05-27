package may2024

import (
	"reflect"
	"testing"
)

func Test_largestLocal(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{
			grid: [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}},
			want: [][]int{{9, 9}, {8, 6}},
		},
		{
			grid: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
			want: [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestLocal(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
