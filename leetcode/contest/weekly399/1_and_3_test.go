package weekly399

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums1: []int{1, 3, 4},
				nums2: []int{1, 3, 4},
				k:     1,
			},
			want: 5,
		},
		{
			args: args{
				nums1: []int{1, 11, 4},
				nums2: []int{1, 3, 4},
				k:     1,
			},
			want: 4,
		},
		{
			args: args{
				nums1: []int{1, 2, 4, 12},
				nums2: []int{2, 4},
				k:     3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				if got := numberOfPairs2(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
					t.Errorf("numberOfPairs2() = %v, want %v", got, tt.want)
				}
			})
			t.Run(tt.name, func(t *testing.T) {
				if got := numberOfPairs1(tt.args.nums1, tt.args.nums2, tt.args.k); int64(got) != tt.want {
					t.Errorf("numberOfPairs1() = %v, want %v", got, tt.want)
				}
			})
			if got := numberOfPairs(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numbersOfPairs(t *testing.B) {
	nums1 := make([]int, 0, 1e4)
	nums2 := make([]int, 0, 1e4)
	for range int(5 * 1e4) {
		nums1 = append(nums1, 4)
		nums1 = append(nums1, 8)
		nums2 = append(nums2, 1)
		nums2 = append(nums2, 2)
	}
	k := 2
	t.Run("O(n*sqrt(m))", func(b *testing.B) {
		got := numberOfPairs2(nums1, nums2, k)
		assert.Equal(b, int64(1e10), got)
	})
	t.Run("O(n*m)", func(b *testing.B) {
		got := numberOfPairs(nums1, nums2, k)
		assert.Equal(b, int64(1e10), got)
	})
}
