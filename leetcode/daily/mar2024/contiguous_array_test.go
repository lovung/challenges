package mar2024

import "testing"

func Test_findMaxLength(t *testing.T) {
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
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 0, 0, 0, 1, 1, 0, 0, 0},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength1(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength2(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength2() = %v, want %v", got, tt.want)
			}
		})
	}
}
