package jul2024

import "testing"

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "aababbab"}, want: 2},
		{args: args{s: "bbaaaaabb"}, want: 2},
		{args: args{s: "a"}, want: 0},
		{args: args{s: "b"}, want: 0},
		{args: args{s: "ba"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
