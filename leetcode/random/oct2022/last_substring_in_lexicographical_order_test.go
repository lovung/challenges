package oct2022

import "testing"

func Test_lastSubstring(t *testing.T) {
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
				s: "abab",
			},
			want: "bab",
		},
		{
			args: args{
				s: "theleetcode",
			},
			want: "theleetcode",
		},
		{
			args: args{
				s: "taleetcode",
			},
			want: "tcode",
		},
		{
			args: args{
				s: "leetcode",
			},
			want: "tcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastSubstring(tt.args.s); got != tt.want {
				t.Errorf("lastSubstring() = %v, want %v", got, tt.want)
			}
			if got := lastSubstring2(tt.args.s); got != tt.want {
				t.Errorf("lastSubstring2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/random/oct2022
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_lastSubstring-12    	  456792	      2193 ns/op	    2240 B/op	      37 allocs/op
func Benchmark_lastSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lastSubstring("leetcode")
		lastSubstring("theleetcode")
		lastSubstring("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		lastSubstring("bababababababababababababababababababababababababababababababa")
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/random/oct2022
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_lastSubstring2-12    	 8519696	       141.1 ns/op	       0 B/op	       0 allocs/op
func Benchmark_lastSubstring2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lastSubstring2("leetcode")
		lastSubstring2("theleetcode")
		lastSubstring2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		lastSubstring2("bababababababababababababababababababababababababababababababa")
	}
}
