package contest

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/longest-increasing-subsequence-ii/
func lengthOfLIS(nums []int, k int) int {
	lis := make([]int, 1e5+1)
	ans := 0
	for i := range nums {
		l := maths.Max(0, nums[i]-k)
		lis[nums[i]] = maths.Max(lis[l:nums[i]]...) + 1
		ans = maths.Max(ans, lis[nums[i]])
	}
	return ans
}

func lengthOfLIS2(nums []int, k int) int {
	ans := 0
	seg := newSegmentTree(1e5 + 1)
	for i := range nums {
		l := maths.Max(0, nums[i]-k)
		r := nums[i]
		res := seg.query(l, r) + 1
		ans = maths.Max(res, ans)
		seg.update(nums[i], res)
	}
	return ans
}

type segmentTree struct {
	data []int
	size int
}

func newSegmentTree(size int) segmentTree {
	return segmentTree{
		data: make([]int, size*2),
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
