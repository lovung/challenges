package oct2022

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 1, 2},
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{1, 2, 4, 5, 7},
			},
			want: 16,
		},
		{
			args: args{
				nums: []int{1, 2, 4, 5, 10, 20},
			},
			want: 11,
		},
		{
			args: args{
				nums: []int{1, 2, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
