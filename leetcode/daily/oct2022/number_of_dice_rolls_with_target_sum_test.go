package oct2022

import "testing"

func Test_numRollsToTarget(t *testing.T) {
	type args struct {
		n      int
		k      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:      1,
				k:      6,
				target: 3,
			},
			want: 1,
		},
		{
			args: args{
				n:      2,
				k:      6,
				target: 7,
			},
			want: 6,
		},
		{
			args: args{
				n:      3,
				k:      6,
				target: 8,
			},
			want: 21,
		},
		{
			args: args{
				n:      30,
				k:      30,
				target: 500,
			},
			want: 222616187,
		},
		{
			args: args{
				n:      20,
				k:      19,
				target: 233,
			},
			want: 378846878,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRollsToTarget(tt.args.n, tt.args.k, tt.args.target); got != tt.want {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := numRollsToTarget2(tt.args.n, tt.args.k, tt.args.target); got != tt.want {
				t.Errorf("numRollsToTarget2() = %v, want %v", got, tt.want)
			}
		})
	}
}
