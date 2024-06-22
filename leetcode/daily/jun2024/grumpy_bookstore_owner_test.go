package jun2024

import "testing"

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
				grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
				minutes:   3,
			},
			want: 16,
		},
		{
			args: args{
				customers: []int{1},
				grumpy:    []int{0},
				minutes:   1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
