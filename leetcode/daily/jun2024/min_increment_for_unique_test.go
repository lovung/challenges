package jun2024

import "testing"

func Test_minIncrementForUnique(t *testing.T) {
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
				nums: []int{1, 2, 2},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{3, 1, 2, 1, 2, 7},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{0, 2, 2},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{0, 2, 2, 3},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{3, 1, 2, 1, 2, 5},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrementForUnique(tt.args.nums); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
