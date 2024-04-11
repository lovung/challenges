package apr2024

import "testing"

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       0,
			},
			want: 8,
		},
		{
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       1,
			},
			want: 2,
		},
		{
			args: args{
				tickets: []int{2, 3, 2},
				k:       2,
			},
			want: 6,
		},
		{
			args: args{
				tickets: []int{2, 3, 2},
				k:       1,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
