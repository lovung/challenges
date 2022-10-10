package oct2022

import "testing"

func Test_breakPalindrome(t *testing.T) {
	type args struct {
		palindrome string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				palindrome: "abccba",
			},
			want: "aaccba",
		},
		{
			args: args{
				palindrome: "aba",
			},
			want: "abb",
		},
		{
			args: args{
				palindrome: "aaaa",
			},
			want: "aaab",
		},
		{
			args: args{
				palindrome: "a",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakPalindrome(tt.args.palindrome); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
