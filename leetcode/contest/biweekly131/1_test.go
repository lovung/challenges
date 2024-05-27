package biweekly131

import "testing"

func Test_duplicateNumbersXOR(t *testing.T) {
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
				nums: []int{1, 2, 1, 3},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1, 2, 2, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duplicateNumbersXOR(tt.args.nums); got != tt.want {
				t.Errorf("duplicateNumbersXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
