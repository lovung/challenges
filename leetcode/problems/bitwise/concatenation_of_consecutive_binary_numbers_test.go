package bitwise

import (
	"testing"
)

func Test_concatenatedBinary(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			args: args{
				n: 3,
			},
			want: 27,
		},
		{
			args: args{
				n: 12,
			},
			want: 505379714,
		},
		{
			args: args{
				n: 18,
			},
			want: 632453782,
		},
		{
			args: args{
				n: 25,
			},
			want: 350431973,
		},
		{
			args: args{
				n: 42,
			},
			want: 727837408,
		},
		{
			args: args{
				n: 98247,
			},
			want: 826411420,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatenatedBinary(tt.args.n); got != tt.want {
				t.Errorf("concatenatedBinary() = %v, want %v", got, tt.want)
			}
			if got := concatenatedBinary2(tt.args.n); got != tt.want {
				t.Errorf("concatenatedBinary2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/grind75
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_concatenatedBinary-12    	      74	  21842979 ns/op	 8395560 B/op	      35 allocs/op
func Benchmark_concatenatedBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenatedBinary(98247)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/grind75
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_concatenatedBinary2-12    	    1976	    616288 ns/op	       0 B/op	       0 allocs/op
func Benchmark_concatenatedBinary2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenatedBinary2(98247)
	}
}
