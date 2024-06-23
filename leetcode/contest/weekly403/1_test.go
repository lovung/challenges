package weekly403

import "testing"

func Test_minimumAverage(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				nums: []int{7, 8, 3, 4, 15, 13, 4, 1},
			},
			want: 5.5,
		},
		{
			args: args{
				nums: []int{1, 9, 8, 3, 10, 5},
			},
			want: 5.5,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 7, 8, 9},
			},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAverage(tt.args.nums); got != tt.want {
				t.Errorf("minimumAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
