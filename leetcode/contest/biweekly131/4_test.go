package biweekly131

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_getResults(t *testing.T) {
	// worstCase1 := make([][]int, 0)
	// const maxVal = int(15 * 1e4)
	// for i := range maxVal / 3 {
	// 	worstCase1 = append(worstCase1, []int{1, maxVal/3 - i})
	// 	worstCase1 = append(worstCase1, []int{2, i, i})
	// }
	type args struct {
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{
				queries: [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}},
			},
			want: []bool{false, true, true},
		},
		{
			args: args{
				queries: [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}},
			},
			want: []bool{true, true, false},
		},
		{
			args: args{
				queries: [][]int{{2, 1, 1}},
			},
			want: []bool{true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResults(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResults() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := getResults1(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResults1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := getResults2(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResults2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_getResults/Segment_Tree:_O(n*log(n))-12         	1000000000	         0.1960 ns/op	       0 B/op	       0 allocs/op
// Benchmark_getResults/Binary_Search_-_heavy_read:_O(n*m)-12         	       1	1846682917 ns/op	 2514944 B/op	       3 allocs/op
// Benchmark_getResults/Binary_Search_-_heavy_write:_O(n*m)-12        	1000000000	         0.2604 ns/op	       0 B/op	       0 allocs/op
func Benchmark_getResults(t *testing.B) {
	worstCase1 := make([][]int, 0)
	const maxVal = int(15 * 1e4)
	for i := range maxVal / 3 {
		worstCase1 = append(worstCase1, []int{1, maxVal/3 - i})
		worstCase1 = append(worstCase1, []int{2, rand.Intn(i + 1), rand.Intn(i + 1)})
	}
	t.Run("Segment Tree: O(n*log(n))", func(t *testing.B) {
		_ = getResults(worstCase1)
	})
	t.Run("Binary Search - heavy read: O(n*m)", func(t *testing.B) {
		_ = getResults1(worstCase1)
	})
	t.Run("Binary Search - heavy write: O(n*m)", func(t *testing.B) {
		_ = getResults2(worstCase1)
	})
}
