package maychallenge

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				image:    [][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}},
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{[]int{2, 2, 2}, []int{2, 2, 0}, []int{2, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floodFillPointer(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor *int
		h        *int
		w        *int
		curColor *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floodFillPointer(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor, tt.args.h, tt.args.w, tt.args.curColor)
		})
	}
}
