package oct2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				envelopes: [][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500},
					{5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}},
			},
			want: 5,
		},
		{
			args: args{
				envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			},
			want: 3,
		},
		{
			args: args{
				envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}},
			},
			want: 1,
		},
		{
			args: args{
				envelopes: [][]int{{46, 89}, {50, 53}, {52, 68}, {72, 45}, {77, 81}},
			},
			want: 3,
		},
		{
			args: args{
				envelopes: [][]int{{1, 15}, {7, 18}, {7, 6}, {7, 100}, {2, 200},
					{17, 30}, {17, 45}, {3, 5}, {7, 8}, {3, 6}, {3, 10}, {7, 20}, {17, 3}, {17, 45}},
			},
			want: 3,
		},
		{
			args: args{
				envelopes: [][]int{{10, 8}, {1, 12}, {6, 15}, {2, 18}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes2(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Using SegmentTree:
// * O(N^2) worst case
// * O(N*logN) best case
// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/random/oct2022
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_maxEnvelopes-12    	    8758	    154674 ns/op	   24740 B/op	       8 allocs/op
func Benchmark_maxEnvelopes(b *testing.B) {
	n := 1000
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		e[i] = []int{i + 1, i + 1}
	}
	for i := 0; i < b.N; i++ {
		assert.Equal(b, n, maxEnvelopes(e))
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/random/oct2022
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_maxEnvelopes2-12    	   68662	     15699 ns/op	    8304 B/op	       6 allocs/op
func Benchmark_maxEnvelopes2(b *testing.B) {
	n := 1000
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		e[i] = []int{i + 1, i + 1}
	}
	for i := 0; i < b.N; i++ {
		assert.Equal(b, n, maxEnvelopes2(e))
	}
}
