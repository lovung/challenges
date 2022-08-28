package medium

import (
	"reflect"
	"testing"
)

func Test_diagonalSort(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				mat: [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}},
			},
			want: [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}},
		},
		{
			name: "test2",
			args: args{
				mat: [][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}},
			},
			want: [][]int{{5, 17, 4, 1, 52, 7}, {11, 11, 25, 45, 8, 69}, {14, 23, 25, 44, 58, 15}, {22, 27, 31, 36, 50, 66}, {84, 28, 75, 33, 55, 68}},
		},
		{
			name: "test3",
			args: args{
				mat: [][]int{{75, 21, 13, 24, 8}, {24, 100, 40, 49, 62}},
			},
			want: [][]int{{75, 21, 13, 24, 8}, {24, 100, 40, 49, 62}},
		},
		{
			name: "test4",
			args: args{
				mat: [][]int{{75, 21, 13, 24, 8}},
			},
			want: [][]int{{75, 21, 13, 24, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalSort(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diagonalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
