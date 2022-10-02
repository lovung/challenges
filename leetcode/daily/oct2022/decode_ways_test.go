package oct2022

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "1201234",
			},
			want: 3,
		},
		{
			args: args{
				s: "12",
			},
			want: 2,
		},
		{
			args: args{
				s: "226",
			},
			want: 3,
		},
		{
			args: args{
				s: "987654321",
			},
			want: 2,
		},
		{
			args: args{
				s: "11106",
			},
			want: 2,
		},
		{
			args: args{
				s: "06",
			},
			want: 0,
		},
		{
			args: args{
				s: "0",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
			if got := numDecodings2(tt.args.s); got != tt.want {
				t.Errorf("numDecodings2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/daily/oct2022
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_numDecodings-12    	       1	11088385336 ns/op	      56 B/op	       2 allocs/op
func Benchmark_numDecodings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numDecodings("111111111111111111111111111111111111111111111")
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/daily/oct2022
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_numDecodings2-12    	 2235020	       551.6 ns/op	     384 B/op	       1 allocs/op
func Benchmark_numDecodings2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numDecodings2("111111111111111111111111111111111111111111111")
	}
}
