package jul2024

import "testing"

func Test_maximumGain(t *testing.T) {
	type args struct {
		s string
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha",
				x: 1926,
				y: 4320,
			},
			want: 112374,
		},
		{
			args: args{
				s: "aabbab",
				x: 1926,
				y: 4320,
			},
			want: 8172,
		},
		{
			args: args{
				s: "bbaaba",
				y: 1926,
				x: 4320,
			},
			want: 8172,
		},
		{
			args: args{
				s: "aabbaaxybbaabb",
				x: 5,
				y: 4,
			},
			want: 20,
		},
		{
			args: args{
				s: "cdbcbbaaabab",
				x: 4,
				y: 5,
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGain(tt.args.s, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("maximumGain() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGain2(tt.args.s, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("maximumGain2() = %v, want %v", got, tt.want)
			}
		})
	}
}
