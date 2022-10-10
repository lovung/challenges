package oct2022

import (
	"testing"
)

func Test_checkPalindromeFormation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				a: "pvhmupgqeltozftlmfjjde",
				b: "yjgpzbezspnnpszebzmhvp",
			},
			want: true,
		},
		{
			args: args{
				a: "yjgpzbezspnnpszebzmhvp",
				b: "pvhmupgqeltozftlmfjjde",
			},
			want: true,
		},
		{
			args: args{
				a: "x",
				b: "y",
			},
			want: true,
		},
		{
			args: args{
				a: "xyz",
				b: "aba",
			},
			want: true,
		},
		{
			args: args{
				a: "xbdef",
				b: "xecab",
			},
			want: false,
		},
		{
			args: args{
				a: "ulacfd",
				b: "jizalu",
			},
			want: true,
		},
		{
			args: args{
				a: "ulaafd",
				b: "jizmlu",
			},
			want: true,
		},
		{
			args: args{
				a: "jizmlu",
				b: "ulaafd",
			},
			want: true,
		},
		{
			args: args{
				a: "abda",
				b: "acmc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindromeFormation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("checkPalindromeFormation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPalindrome2NonPalindrome(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindrome2NonPalindrome(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("checkPalindrome2NonPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
