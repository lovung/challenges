package apr2024

import "testing"

func Test_longestIdealString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "acfgbd",
				k: 2,
			},
			want: 4,
		},
		{
			args: args{
				s: "acbd",
				k: 3,
			},
			want: 4,
		},
		{
			args: args{
				s: "azaza",
				k: 25,
			},
			want: 5,
		},
		{
			args: args{
				s: "acfgbdbvxchdasx",
				k: 4,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIdealString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestIdealString() = %v, want %v", got, tt.want)
			}
		})
	}
}
