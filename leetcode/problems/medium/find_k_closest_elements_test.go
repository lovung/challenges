package medium

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr: []int{1, 5, 10},
				k:   1,
				x:   3,
			},
			want: []int{1},
		},
		{
			args: args{
				arr: []int{1, 5, 10},
				k:   1,
				x:   4,
			},
			want: []int{5},
		},
		{
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   10,
			},
			want: []int{2, 3, 4, 5},
		},
		{
			args: args{
				arr: []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
				k:   3,
				x:   5,
			},
			want: []int{3, 3, 4},
		},
		{
			args: args{
				arr: []int{0, 2, 2, 3, 4, 6, 7, 8, 9, 9},
				k:   4,
				x:   5,
			},
			want: []int{3, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
			if got := findClosestElements2(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   0,
			},
			want:  -1,
			want1: 0,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   1,
			},
			want:  0,
			want1: 1,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   2,
			},
			want:  0,
			want1: 1,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   3,
			},
			want:  1,
			want1: 2,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   4,
			},
			want:  1,
			want1: 2,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   5,
			},
			want:  2,
			want1: 3,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   6,
			},
			want:  2,
			want1: 3,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   7,
			},
			want:  3,
			want1: 4,
		},
		{
			args: args{
				arr: []int{1, 3, 5, 7},
				x:   8,
			},
			want:  3,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := binarySearch(tt.args.arr, tt.args.x)
			if got != tt.want {
				t.Errorf("binarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("binarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
