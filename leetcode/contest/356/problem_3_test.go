package contest356

import (
	"testing"
)

func Test_cntMergeLen(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{"ab", "aba"},
			want: 2,
		},
		{
			args: args{"ab", "bbba"},
			want: 1,
		},
		{
			args: args{"abb", "bbba"},
			want: 2,
		},
		{
			args: args{"abc", "aaa"},
			want: 0,
		},
		{
			args: args{"abc", "bca"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cntMergeLen(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("cntMergeLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tryMerge(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{"ab", "bbba"},
			want: "abbba",
		},
		{
			args: args{"abb", "bbba"},
			want: "abbba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tryMerge(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("tryMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumString(t *testing.T) {
	type args struct {
		a string
		b string
		c string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{"cab", "a", "b"},
			want: "cab",
		},
		{
			args: args{"abc", "bca", "aaa"},
			want: "aaabca",
		},
		{
			args: args{"ab", "ba", "aba"},
			want: "aba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumString(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("minimumString() = %v, want %v", got, tt.want)
			}
		})
	}
}
