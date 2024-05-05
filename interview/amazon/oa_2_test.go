package amazon

import (
	"math/rand"
	"testing"
)

func Test_getMaximumPoints(t *testing.T) {
	type args struct {
		days []int32
		k    int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				days: []int32{1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			args: args{
				days: []int32{7, 4, 6, 3, 5},
				k:    8,
			},
			want: 33,
		},
		{
			args: args{
				days: []int32{2, 3, 2},
				k:    4,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumPoints1(tt.args.days, tt.args.k); got != tt.want {
				t.Errorf("getMaximumPoints1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumPoints2(tt.args.days, tt.args.k); got != tt.want {
				t.Errorf("getMaximumPoints2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumPoints3(tt.args.days, tt.args.k); got != tt.want {
				t.Errorf("getMaximumPoints3() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumPoints4(tt.args.days, tt.args.k); got != tt.want {
				t.Errorf("getMaximumPoints4() = %v, want %v", got, tt.want)
			}
		})
		// t.Run(tt.name, func(t *testing.T) {
		// 	if got := getMaximumPoints5(tt.args.days, tt.args.k); got != tt.want {
		// 		t.Errorf("getMaximumPoints4() = %v, want %v", got, tt.want)
		// 	}
		// })
	}
}

// Benchmark_getMaximumPoints/1-12  	1000000000	         0.0004313 ns/op	       0 B/op	       0 allocs/op
// Benchmark_getMaximumPoints/2-12  	1000000000	         0.1319 ns/op	       0 B/op	       0 allocs/op
// Benchmark_getMaximumPoints/3-12  	1000000000	         0.04173 ns/op	       0 B/op	       0 allocs/op
// Benchmark_getMaximumPoints/4-12  	1000000000	         0.0008611 ns/op	       0 B/op	       0 allocs/op
func Benchmark_getMaximumPoints(b *testing.B) {
	genDays := make([]int32, 1000)
	k := int32(rand.Intn(1000))
	for i := 0; i < 1000; i++ {
		genDays[i] = int32(rand.Intn(1000))
	}

	for range b.N {
		b.ResetTimer()
		b.Run("1", func(b *testing.B) {
			getMaximumPoints1(genDays, k)
		})
		b.ResetTimer()
		b.Run("2", func(b *testing.B) {
			getMaximumPoints2(genDays, k)
		})
		b.ResetTimer()
		b.Run("3", func(b *testing.B) {
			getMaximumPoints3(genDays, k)
		})
		b.ResetTimer()
		b.Run("4", func(b *testing.B) {
			getMaximumPoints4(genDays, k)
		})
	}
}

// 1e4
// Benchmark_getMaximumPoints2/1-12  	1000000000	         0.04242 ns/op	       0 B/op	       0 allocs/op
// Benchmark_getMaximumPoints2/4-12  	1000000000	         0.09175 ns/op	       0 B/op	       0 allocs/op

// 1e5
// Benchmark_getMaximumPoints2/1-12  	       1	38911274250 ns/op	19968024192 B/op	      16 allocs/op
// Benchmark_getMaximumPoints2/4-12  	       1	9157720625 ns/op	  802816 B/op	       1 allocs/op
func Benchmark_getMaximumPoints2(b *testing.B) {
	k := int32(rand.Intn(1e5))
	genDays := make([]int32, 1e5)
	for i := 0; i < 1e5; i++ {
		genDays[i] = int32(rand.Intn(1e5))
	}

	for range b.N {
		b.ResetTimer()
		b.Run("1", func(b *testing.B) {
			getMaximumPoints1(genDays, k)
		})
		b.ResetTimer()
		b.Run("4", func(b *testing.B) {
			getMaximumPoints4(genDays, k)
		})
	}
}

func Test_calculatePointAtDay(t *testing.T) {
	type args struct {
		days   []int32
		dayIdx int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 0,
			},
			want: 1,
		},
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 1,
			},
			want: 2,
		},
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 2,
			},
			want: 1,
		},
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 3,
			},
			want: 2,
		},
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 4,
			},
			want: 3,
		},
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 5,
			},
			want: 1,
		},
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 6,
			},
			want: 2,
		},
		{
			args: args{
				days:   []int32{2, 3, 2},
				dayIdx: 7,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		// t.Run(tt.name, func(t *testing.T) {
		// 	cache := make(map[int]int32)
		// 	if got := calculatePointAtDay(cache, tt.args.days, tt.args.dayIdx); got != tt.want {
		// 		t.Errorf("calculatePointAtDay() = %v, want %v", got, tt.want)
		// 	}
		// })
		t.Run(tt.name, func(t *testing.T) {
			cache := make(map[int]int32)
			prefix := prefixSum(tt.args.days)
			if got := calculatePointAtDayWithPrefixSum(cache, prefix, tt.args.dayIdx); got != tt.want {
				t.Errorf("calculatePointAtDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
