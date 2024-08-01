package jul2024

import (
	"reflect"
	"testing"
)

func Test_restoreMatrix(t *testing.T) {
	type args struct {
		rowSum []int
		colSum []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				rowSum: []int{3, 8},
				colSum: []int{4, 7},
			},
			want: [][]int{{3, 0}, {1, 7}},
		},
		{
			args: args{
				rowSum: []int{5, 7, 10},
				colSum: []int{8, 6, 8},
			},
			want: [][]int{{5, 0, 0}, {3, 4, 0}, {0, 2, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreMatrix(tt.args.rowSum, tt.args.colSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
