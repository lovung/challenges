package jul2023

import (
	"math"
	"testing"
)

func Test_soupServings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{n: 50},
			want: 0.62500,
		},
		{
			args: args{n: 100},
			want: 0.71875,
		},
		{
			args: args{n: 1000},
			want: 0.97657,
		},
		{
			args: args{n: 660295675},
			want: 1.00000,
		},
		{
			args: args{n: 80},
			want: 0.71875,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soupServings(tt.args.n); math.Round(got*1e5)/1e5 != tt.want {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}
