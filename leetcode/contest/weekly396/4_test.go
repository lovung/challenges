package weekly396

import "testing"

func Test_minCostToEqualizeArray(t *testing.T) {
	type args struct {
		nums  []int
		cost1 int
		cost2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:  []int{1, 1000000, 999999},
				cost1: 1e6,
				cost2: 1,
			},
			want: 1999997,
		},
		{
			args: args{
				nums:  []int{1000000, 1000000, 7, 1000000, 1000000},
				cost1: 1e6,
				cost2: 1,
			},
			want: 1333324,
		},
		{
			args: args{
				nums:  []int{1, 14, 15},
				cost1: 2,
				cost2: 1,
			},
			want: 27,
		},
		{
			args: args{
				nums:  []int{15},
				cost1: 2,
				cost2: 1,
			},
			want: 0,
		},
		{
			args: args{
				nums:  []int{40, 23, 35},
				cost1: 14,
				cost2: 2,
			},
			want: 58,
		},
		{
			args: args{
				nums:  []int{1, 1000000},
				cost1: 1000000,
				cost2: 1,
			},
			want: 998993007,
		},
		{
			args: args{
				nums:  []int{4, 3},
				cost1: 2,
				cost2: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostToEqualizeArray(tt.args.nums, tt.args.cost1, tt.args.cost2); got != tt.want {
				t.Errorf("minCostToEqualizeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
