package litmus

import "testing"

func Test_findMaxSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{2, 1}},
			want: 3,
		},
		{
			args: args{[]int{5, 7, 9, 11}},
			want: 20,
		},
		{
			args: args{[]int{5, 7, 9, 11, 11}},
			want: 22,
		},
		{
			args: args{[]int{1, 1, 2, 10, 2, 3, 10, 11}},
			want: 21,
		},
		{
			args: args{[]int{1, 1, 2, 10, 2, 3, 10}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxSum(tt.args.numbers); got != tt.want {
				t.Errorf("findMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
