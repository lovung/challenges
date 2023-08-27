package contest

import "testing"

func Test_minimumSeconds(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{8, 8, 9, 10, 9}},
			want: 1,
		},
		{
			args: args{nums: []int{1, 2, 1, 2}},
			want: 1,
		},
		{
			args: args{nums: []int{2, 1, 3, 3, 2}},
			want: 2,
		},
		{
			args: args{nums: []int{5, 5, 5, 5}},
			want: 0,
		},
		{
			args: args{nums: []int{1, 11, 11, 11, 19, 12, 8, 7, 19}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSeconds(tt.args.nums); got != tt.want {
				t.Errorf("minimumSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
