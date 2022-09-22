package containsduplicatie

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/contains-duplicate-iii/
const offset = 1e9 + 1
const idxOffset = 1

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	lastIdx := newSegmentTree(2*1e9 + 1)
	for i := range nums {
		if idx := lastIdx.query(
			nums[i]+offset-valueDiff,
			nums[i]+offset+valueDiff+1,
		); idx > 0 && i-(idx-idxOffset) <= indexDiff {
			return true
		}
		lastIdx.update(nums[i]+offset, i+idxOffset)
	}
	return false
}

type segmentTree struct {
	data map[int]int
	size int
}

func newSegmentTree(size int) segmentTree {
	return segmentTree{
		data: make(map[int]int),
		size: size,
	}
}

func (st *segmentTree) query(lo, hi int) int {
	lo += st.size
	hi += st.size

	res := 0
	for lo < hi {
		if lo&1 != 0 {
			res = maths.Max(res, st.data[lo])
			lo++
		}
		if hi&1 != 0 {
			hi--
			res = maths.Max(res, st.data[hi])
		}
		lo >>= 1
		hi >>= 1
	}
	return res
}

func (st *segmentTree) update(pos, val int) {
	pos += st.size
	st.data[pos] = val

	for pos > 1 {
		pos >>= 1
		st.data[pos] = maths.Max(st.data[2*pos], st.data[2*pos+1])
	}
}
