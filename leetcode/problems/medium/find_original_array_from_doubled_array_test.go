package medium

import (
	"reflect"
	"testing"
)

func Test_findOriginalArray(t *testing.T) {
	type args struct {
		changed []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				changed: []int{2, 2, 1, 1},
			},
			want: []int{1, 1},
		},
		{
			args: args{
				changed: []int{1, 3, 4, 2, 6, 8},
			},
			want: []int{1, 3, 4},
		},
		{
			args: args{
				changed: []int{6, 3, 0, 1},
			},
			want: []int{},
		},
		{
			args: args{
				changed: []int{1},
			},
			want: []int{},
		},
		{
			args: args{
				changed: []int{4, 2, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			args: args{
				changed: []int{5, 8, 7, 8, 16, 5, 16, 14, 10, 10},
			},
			want: []int{5, 5, 7, 8, 8},
		},
		{
			args: args{
				changed: []int{4, 4, 16, 20, 8, 8, 2, 10},
			},
			want: []int{2, 4, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOriginalArray(tt.args.changed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOriginalArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
