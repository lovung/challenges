package grind75

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				intervals: [][]int{
					{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {17, 18},
				},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}, {17, 18}},
		},
		{
			args: args{
				intervals: [][]int{
					{1, 3}, {6, 9},
				},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			args: args{
				intervals: [][]int{
					{1, 3}, {6, 9},
				},
				newInterval: []int{1, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			args: args{
				intervals: [][]int{
					{1, 3}, {6, 9},
				},
				newInterval: []int{10, 11},
			},
			want: [][]int{{1, 3}, {6, 9}, {10, 11}},
		},
		{
			args: args{
				intervals: [][]int{
					{6, 9}, {10, 11},
				},
				newInterval: []int{1, 3},
			},
			want: [][]int{{1, 3}, {6, 9}, {10, 11}},
		},
		{
			args: args{
				intervals:   [][]int{{1, 5}, {6, 9}},
				newInterval: []int{5, 6},
			},
			want: [][]int{{1, 9}},
		},
		{
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			args: args{
				intervals:   [][]int{},
				newInterval: []int{6, 8},
			},
			want: [][]int{{6, 8}},
		},
		{
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{6, 8},
			},
			want: [][]int{{1, 5}, {6, 8}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("sol_1_case_%d", i), func(t *testing.T) {
			interval := slices.Clone(tt.args.intervals)
			newInterval := slices.Clone(tt.args.newInterval)
			if got := insert(interval, newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("sol_2_case_%d", i), func(t *testing.T) {
			interval := slices.Clone(tt.args.intervals)
			newInterval := slices.Clone(tt.args.newInterval)
			if got := insert2(interval, newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("sol_3_case_%d", i), func(t *testing.T) {
			interval := slices.Clone(tt.args.intervals)
			newInterval := slices.Clone(tt.args.newInterval)
			if got := insert3(interval, newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_insert(b *testing.B) {
	init := make([][]int, 0, int(1e6))
	for i := range int(1e6) {
		init = append(init, []int{2 * i, 2*i + 1})
	}

	b.ResetTimer()
	bigIntervals := slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_1_case_1", func(b *testing.B) {
		for range b.N {
			insert(bigIntervals, []int{1, 100})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_1_case_100", func(b *testing.B) {
		for range b.N {
			insert(bigIntervals, []int{100, 200})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_1_case_1000", func(b *testing.B) {
		for range b.N {
			insert(bigIntervals, []int{1000, 1010})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_1_case_10000", func(b *testing.B) {
		for range b.N {
			insert(bigIntervals, []int{10000, 10010})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_1_case_100000", func(b *testing.B) {
		for range b.N {
			insert(bigIntervals, []int{100000, 100100})
		}
	})

	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_2_case_1", func(b *testing.B) {
		for range b.N {
			insert2(bigIntervals, []int{1, 100})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_2_case_100", func(b *testing.B) {
		for range b.N {
			insert2(bigIntervals, []int{100, 200})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_2_case_1000", func(b *testing.B) {
		for range b.N {
			insert2(bigIntervals, []int{1000, 1010})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_2_case_10000", func(b *testing.B) {
		for range b.N {
			insert2(bigIntervals, []int{10000, 10010})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_2_case_100000", func(b *testing.B) {
		for range b.N {
			insert2(bigIntervals, []int{100000, 100100})
		}
	})

	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_3_case_1", func(b *testing.B) {
		for range b.N {
			insert3(bigIntervals, []int{1, 100})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_3_case_100", func(b *testing.B) {
		for range b.N {
			insert3(bigIntervals, []int{100, 200})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_3_case_1000", func(b *testing.B) {
		for range b.N {
			insert3(bigIntervals, []int{1000, 1010})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_3_case_10000", func(b *testing.B) {
		for range b.N {
			insert3(bigIntervals, []int{10000, 10010})
		}
	})
	b.ResetTimer()
	bigIntervals = slices.Clone(init)
	b.ResetTimer()
	b.Run("sol_3_case_100000", func(b *testing.B) {
		for range b.N {
			insert3(bigIntervals, []int{100000, 100100})
		}
	})
}
