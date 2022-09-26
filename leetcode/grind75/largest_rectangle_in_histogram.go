package grind75

import (
	"math"

	"github.com/lovung/ds/maths"
	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	// max of all subArray ( 					-> O (N^2)		* problem here
	// 		min in subArray (min_height)		-> segmentTree (O(logN))
	// 			x
	// 		size of subArray					-> O(1)
	// )										-> O (N^2 * logN)
	n := len(heights)
	minFunc := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	segTree := trees.NewSegmentTreeWithArray(n, minFunc)
	segTree.SetInitQueryValue(math.MaxInt64)
	maxArea := 0
	for i, h := range heights {
		segTree.Update(i, h)
		maxArea = maths.Max(maxArea, h)
	}
	for i, h := range heights {
		if h == 0 {
			continue
		}
		minWidth := maxArea / h
		for j := n - 1; j >= i+minWidth; j-- {
			size := j - i + 1
			minHeight := segTree.Query(i, j+1)
			area := size * minHeight
			maxArea = maths.Max(maxArea, area)
		}
	}
	return maxArea
}

type item struct {
	val int
	idx int
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	minFunc := func(a, b item) item {
		if a.val < b.val {
			return a
		}
		return b
	}
	segTree := trees.NewSegmentTreeWithArray(n, minFunc)
	segTree.SetInitQueryValue(item{val: math.MaxInt64, idx: -1})
	maxArea := 0
	for i, h := range heights {
		segTree.Update(i, item{val: h, idx: i})
		maxArea = maths.Max(maxArea, h)
	}
	minHeight := segTree.Query(0, n)
	return maths.Max(
		maxArea,
		n*minHeight.val,
		divideToSolve(segTree, heights, 0, minHeight.idx-1),
		divideToSolve(segTree, heights, minHeight.idx+1, n-1),
	)
}

func divideToSolve(segTree trees.SegmentTree[item], heights []int, l, r int) int {
	if l > r {
		return 0
	}
	if l == r {
		return heights[l]
	}
	size := r - l + 1
	minHeight := segTree.Query(l, r+1)
	return maths.Max(
		size*minHeight.val,
		divideToSolve(segTree, heights, l, minHeight.idx-1),
		divideToSolve(segTree, heights, minHeight.idx+1, r),
	)
}
