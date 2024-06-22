package jun2024

import "testing"

func Test_maxProfitAssignment(t *testing.T) {
	type args struct {
		difficulty []int
		profit     []int
		worker     []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				difficulty: []int{23, 30, 35, 35, 43, 46, 47, 81, 83, 98},
				profit:     []int{8, 11, 11, 20, 33, 37, 60, 72, 87, 95},
				worker:     []int{95, 46, 47, 97, 11, 35, 99, 56, 41, 92},
			},
			want: 553,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitAssignment(tt.args.difficulty, tt.args.profit, tt.args.worker); got != tt.want {
				t.Errorf("maxProfitAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
