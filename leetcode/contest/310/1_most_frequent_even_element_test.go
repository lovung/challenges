package contest

import "testing"

func Test_mostFrequentEven(t *testing.T) {
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
				nums: []int{0, 1, 2, 2, 4, 4, 1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{4, 4, 4, 9, 2, 4},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{29, 47, 21, 41, 13, 37, 25, 7},
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{1e5},
			},
			want: 1e5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFrequentEven(tt.args.nums); got != tt.want {
				t.Errorf("mostFrequentEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
