package grind75

import "testing"

func Test_jobScheduling(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		profit    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				startTime: []int{1, 2, 3, 3},
				endTime:   []int{3, 4, 5, 6},
				profit:    []int{50, 10, 40, 70},
			},
			want: 120,
		},
		{
			args: args{
				startTime: []int{1, 2, 3, 4, 6},
				endTime:   []int{3, 5, 10, 6, 9},
				profit:    []int{20, 20, 100, 70, 60},
			},
			want: 150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jobScheduling(tt.args.startTime, tt.args.endTime, tt.args.profit); got != tt.want {
				t.Errorf("jobScheduling() = %v, want %v", got, tt.want)
			}
		})
	}
}
