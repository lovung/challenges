package grind75

// Link: https://leetcode.com/problems/find-median-from-data-stream/
type MedianFinder struct {
	arr []int
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		arr: make([]int, 0, 5*1e4),
	}
}

func (f *MedianFinder) len() int {
	return len(f.arr)
}

// Example cases:
// 1, 3, 5, 7
// 0 -> l = 0, r = -1		want: -1
// 1 -> l = 0, r = 0		want: 0
// 2 -> l = 1, r = 0		want: 0
// 3 -> mid = 1				want: 1
// 7 -> l = 3, r = 3		want: 3
// 8 -> l = 4, r = 3		want: 4
func (f *MedianFinder) binarySearchFindInsertIndex(num int) int {
	n := f.len()
	if n == 0 {
		return -1
	}
	if num < f.arr[0] {
		return -1
	}
	if num > f.arr[n-1] {
		return n
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) >> 1 // % 2
		if f.arr[mid] == num {
			return mid
		}
		if f.arr[mid] < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func (f *MedianFinder) findMidLeftIndex() int {
	return (f.len()+1)>>1 - 1 // (n + 1) / 2 -1
}

func (f *MedianFinder) findMidRightIndex() int {
	return f.len() >> 1
}

func (f *MedianFinder) AddNum(num int) {
	switch insertAfterIndex := f.binarySearchFindInsertIndex(num); insertAfterIndex {
	case f.len():
		f.arr = append(f.arr, num)
	case -1:
		f.arr = append(f.arr, 0)
		copy(f.arr[1:], f.arr)
		f.arr[0] = num
	default:
		f.arr = append(f.arr, 0)
		copy(f.arr[insertAfterIndex+2:], f.arr[insertAfterIndex+1:])
		f.arr[insertAfterIndex+1] = num
	}
}

func (f *MedianFinder) FindMedian() float64 {
	// even length
	if f.len()&1 != 0 {
		return float64(f.arr[f.findMidLeftIndex()])
	}
	// odd length
	return float64(f.arr[f.findMidLeftIndex()]+
		f.arr[f.findMidRightIndex()]) / 2
}
