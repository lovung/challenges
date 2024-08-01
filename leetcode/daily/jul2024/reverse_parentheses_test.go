package jul2024

import "testing"

func Test_reverseParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "(abcd)",
			},
			want: "dcba",
		},
		{
			args: args{
				s: "e(abcd)f",
			},
			want: "edcbaf",
		},
		{
			args: args{
				s: "(u(love)i)",
			},
			want: "iloveu",
		},
		{
			args: args{
				s: "(ed(et(oc))el)",
			},
			want: "leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
