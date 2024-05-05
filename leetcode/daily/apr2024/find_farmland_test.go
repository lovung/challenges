package apr2024

import (
	"reflect"
	"testing"
)

func Test_findFarmland(t *testing.T) {
	tests := []struct {
		name string
		land [][]int
		want [][]int
	}{
		{
			land: [][]int{
				{1, 0, 0},
				{0, 1, 1},
				{0, 1, 1},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{1, 1, 2, 2},
			},
		},
		{
			land: [][]int{
				{1, 1},
				{1, 1},
			},
			want: [][]int{
				{0, 0, 1, 1},
			},
		},
		{
			land: [][]int{
				{1, 0, 0, 0},
				{0, 1, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 0},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{1, 1, 2, 1},
				{1, 3, 2, 3},
				{3, 2, 3, 2},
			},
		},
		{
			land: [][]int{
				{1, 0, 0},
				{0, 1, 1},
				{0, 1, 1},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{1, 1, 2, 2},
			},
		},
		{
			land: [][]int{
				{0},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFarmland(tt.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFarmland() = %v, want %v", got, tt.want)
			}
		})
	}
}
