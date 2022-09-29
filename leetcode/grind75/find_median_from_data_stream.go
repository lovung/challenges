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

func (this *MedianFinder) len() int {
	return len(this.arr)
}

// Example cases:
// 1, 3, 5, 7
// 0 -> l = 0, r = -1		want: -1
// 1 -> l = 0, r = 0		want: 0
// 2 -> l = 1, r = 0		want: 0
// 3 -> mid = 1				want: 1
// 7 -> l = 3, r = 3		want: 3
// 8 -> l = 4, r = 3		want: 4
func (this *MedianFinder) binarySearchFindInsertIndex(num int) int {
	n := this.len()
	if n == 0 {
		return -1
	}
	if num < this.arr[0] {
		return -1
	}
	if num > this.arr[n-1] {
		return n
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) >> 1 // % 2
		if this.arr[mid] == num {
			return mid
		}
		if this.arr[mid] < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func (this *MedianFinder) findMidLeftIndex() int {
	return (this.len()+1)>>1 - 1 // (n + 1) / 2 -1
}

func (this *MedianFinder) findMidRightIndex() int {
	return this.len() >> 1
}

func (this *MedianFinder) AddNum(num int) {
	switch insertAfterIndex := this.binarySearchFindInsertIndex(num); insertAfterIndex {
	case this.len():
		this.arr = append(this.arr, num)
	case -1:
		this.arr = append(this.arr, 0)
		copy(this.arr[1:], this.arr)
		this.arr[0] = num
	default:
		this.arr = append(this.arr, 0)
		copy(this.arr[insertAfterIndex+2:], this.arr[insertAfterIndex+1:])
		this.arr[insertAfterIndex+1] = num
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// even length
	if this.len()&1 != 0 {
		return float64(this.arr[this.findMidLeftIndex()])
	}
	// odd length
	return float64(this.arr[this.findMidLeftIndex()]+
		this.arr[this.findMidRightIndex()]) / 2
}
