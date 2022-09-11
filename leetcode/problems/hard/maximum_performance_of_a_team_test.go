package hard

import "testing"

func Test_maxPerformance(t *testing.T) {
	type args struct {
		n          int
		speed      []int
		efficiency []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:          6,
				speed:      []int{2, 10, 3, 1, 5, 8},
				efficiency: []int{5, 4, 3, 9, 7, 2},
				k:          2,
			},
			want: 60,
		},
		{
			args: args{
				n:          6,
				speed:      []int{2, 10, 3, 1, 5, 8},
				efficiency: []int{5, 4, 3, 9, 7, 2},
				k:          3,
			},
			want: 68,
		},
		{
			args: args{
				n:          6,
				speed:      []int{2, 10, 3, 1, 5, 8},
				efficiency: []int{5, 4, 3, 9, 7, 2},
				k:          4,
			},
			want: 72,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPerformance(tt.args.n, tt.args.speed, tt.args.efficiency, tt.args.k); got != tt.want {
				t.Errorf("maxPerformance() = %v, want %v", got, tt.want)
			}
		})
	}
}
