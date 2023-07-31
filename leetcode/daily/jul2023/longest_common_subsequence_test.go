package jul2023

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s1: "let",
				s2: "leet",
			},
			want: 3,
		},
		{
			args: args{
				s1: "acb",
				s2: "accbxxxd",
			},
			want: 3,
		},
		{
			args: args{
				s1: "acb",
				s2: "accdxxx",
			},
			want: 2,
		},
		{
			args: args{
				s1: "lete",
				s2: "leet",
			},
			want: 3,
		},
		{
			args: args{
				s1: "bbccacacaab",
				s2: "aabccb",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
