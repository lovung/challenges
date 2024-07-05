package weekly404

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    3,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{1, 4, 2, 3, 1, 4},
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
