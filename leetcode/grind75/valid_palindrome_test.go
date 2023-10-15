package grind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			args: args{
				s: " ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.args.s))
			assert.Equal(t, tt.want, isPalindrome2(tt.args.s))
		})
	}
}
