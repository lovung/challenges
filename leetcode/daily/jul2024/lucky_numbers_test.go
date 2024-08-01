package jul2024

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			},
			want: []int{15},
		},
		{
			args: args{
				matrix: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}},
			},
			want: []int{12},
		},
		{
			args: args{
				matrix: [][]int{{7, 8}, {1, 2}},
			},
			want: []int{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("luckyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
