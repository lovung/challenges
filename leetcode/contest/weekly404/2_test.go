package weekly404

import "testing"

func Test_maximumLength1(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{1, 2, 1, 1, 2, 1, 2},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength1(tt.args.nums); got != tt.want {
				t.Errorf("maximumLength1() = %v, want %v", got, tt.want)
			}
		})
	}
}
