package contest356

import "testing"

func Test_countCompleteSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{1, 3, 1, 2, 2}},
			want: 4,
		},
		{
			args: args{[]int{5, 5, 5, 5}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
