package contest

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{2, 1, 3, 2, 1}},
			want: 3,
		},
		{
			args: args{nums: []int{1, 3, 2, 1, 3, 3}},
			want: 2,
		},
		{
			args: args{nums: []int{2, 2, 2, 2, 3, 3}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
