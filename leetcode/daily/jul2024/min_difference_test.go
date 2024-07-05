package jul2024

import "testing"

func Test_minDifference(t *testing.T) {
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
				nums: []int{5, 3, 2, 4},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1, 5, 0, 10, 14},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{3, 100, 20},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifference(tt.args.nums); got != tt.want {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
