package jul2023

import "testing"

func Test_strangePrinter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{"aaabbb"},
			want: 2,
		},
		{
			args: args{"aba"},
			want: 2,
		},
		{
			args: args{"abacaba"},
			want: 4,
		},
		{
			args: args{"tbgtgb"},
			want: 4,
		},
		{
			args: args{"xabcdddddddcba"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.args.s); got != tt.want {
				t.Errorf("strangePrinter(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
