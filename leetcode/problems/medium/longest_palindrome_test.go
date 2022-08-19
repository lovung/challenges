package medium

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "test2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "test3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "test4",
			args: args{
				s: "aaaaabbaa",
			},
			want: "aabbaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/problems/medium
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// Benchmark_longestPalindrome-8   	 1263055	       832.3 ns/op	     504 B/op	      16 allocs/op
// PASS
// ok  	github.com/lovung/challenges/leetcode/problems/medium	2.112s
func Benchmark_longestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome("babad")
		longestPalindrome("aaaaabbaa")
		longestPalindrome("cbbd")
		longestPalindrome("a")
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/problems/medium
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// Benchmark_longestPalindrome2-8   	 1562270	       765.4 ns/op	     504 B/op	      16 allocs/op
// PASS
// ok  	github.com/lovung/challenges/leetcode/problems/medium	2.074s
func Benchmark_longestPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome2("babad")
		longestPalindrome2("aaaaabbaa")
		longestPalindrome2("cbbd")
		longestPalindrome2("a")
	}
}
