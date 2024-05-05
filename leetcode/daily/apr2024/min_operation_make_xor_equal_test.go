package apr2024

import "testing"

func Test_minOperations(t *testing.T) {
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
				nums: []int{2, 1, 3, 4},
				k:    1,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{2, 0, 2, 0},
				k:    0,
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{9, 7, 9, 14, 8, 6},
				k:    12,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
