package grind75

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			args: args{
				s: "ADOBECODEBANC",
				t: "ADOBECODEBANC",
			},
			want: "ADOBECODEBANC",
		},
		{
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			args: args{
				s: "aa",
				t: "a",
			},
			want: "a",
		},
		{
			args: args{
				s: "cgklivwehljxrdzpfdqsapogwvjtvbzahjnsejwnuhmomlfsrvmrnczjzjevkdvroiluthhpqtffhlzyglrvorgnalk",
				t: "mqfff",
			},
			want: "fsrvmrnczjzjevkdvroiluthhpqtff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
			if got := minWindow2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/grind75
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_minWindow-12    	   19648	     65474 ns/op	       0 B/op	       0 allocs/op
func Benchmark_minWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minWindow("cgklivwehljxrdzpfdqsapogwvjtvbzahjnsejwnuhmomlfsrvmrnczjzjevkdvroiluthhpqtffhlzyglrvorgnalk", "mqfff")
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/grind75
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_minWindow2-12    	  247420	      4219 ns/op	       0 B/op	       0 allocs/op
func Benchmark_minWindow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minWindow2("cgklivwehljxrdzpfdqsapogwvjtvbzahjnsejwnuhmomlfsrvmrnczjzjevkdvroiluthhpqtffhlzyglrvorgnalk", "mqfff")
	}
}
