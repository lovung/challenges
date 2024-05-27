package may2024

import "testing"

func Test_specialArray(t *testing.T) {
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
				nums: []int{3, 5},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{0, 3, 0, 4, 4},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{0, 0},
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{3, 6, 7, 7, 0},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialArray(tt.args.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
