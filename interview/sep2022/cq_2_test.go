package sep2022

import "testing"

func Test_findTotalImbalance(t *testing.T) {
	type args struct {
		rank []int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				rank: []int32{4, 1, 3, 2},
			},
			want: 3,
		},
		{
			args: args{
				rank: []int32{1, 2},
			},
			want: 0,
		},
		{
			args: args{
				rank: []int32{5, 1, 3, 4, 2},
			},
			want: 7,
		},
		{
			args: args{
				rank: []int32{1, 5, 3, 4, 2},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTotalImbalance(tt.args.rank); got != tt.want {
				t.Errorf("findTotalImbalance() = %v, want %v", got, tt.want)
			}
			if got := findTotalImbalance2(tt.args.rank); got != tt.want {
				t.Errorf("findTotalImbalance2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/interview/amz
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// Benchmark_findTotalImbalance2-8   	  453951	      2738 ns/op	    1104 B/op	      48 allocs/op
// PASS
// ok  	github.com/lovung/challenges/interview/amz	1.616s
func Benchmark_findTotalImbalance2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findTotalImbalance2([]int32{5, 1, 3, 4, 2})
		findTotalImbalance2([]int32{4, 1, 3, 2})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/interview/sep2022
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// Benchmark_findTotalImbalance-8   	 4728900	       274.6 ns/op	      72 B/op	       9 allocs/op
// PASS
// ok  	github.com/lovung/challenges/interview/sep2022	2.752s
func Benchmark_findTotalImbalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findTotalImbalance([]int32{5, 1, 3, 4, 2})
		findTotalImbalance([]int32{4, 1, 3, 2})
	}
}
