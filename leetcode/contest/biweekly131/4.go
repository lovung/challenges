package biweekly131

import (
	"slices"
)

func getResults1(queries [][]int) []bool {
	obs := make([]int, 0, 15*1e4)
	gaps := make([]int, 0, 15*1e4)
	res := make([]bool, 0, len(queries))
	for _, q := range queries {
		idx, _ := slices.BinarySearch(obs, q[1])
		switch q[0] {
		case 1: // place obs
			obs = slices.Insert(obs, idx, q[1])
			gap := q[1]
			if idx > 0 {
				gap = max(gaps[idx-1], q[1]-obs[idx-1])
			}
			gaps = slices.Insert(gaps, idx, gap)
			for j := idx + 1; j < len(gaps); j++ {
				gaps[j] = max(gaps[j-1], obs[j]-obs[j-1])
			}
		case 2:
			maxGap := q[1]
			if idx > 0 {
				maxGap = max(gaps[idx-1], q[1]-obs[idx-1])
			}
			res = append(res, maxGap >= q[2])
		}
	}
	return res
}

// O(n*m)
func getResults2(queries [][]int) []bool {
	obs := make([]int, 0, 1<<5)
	res := make([]bool, 0, len(queries))
outter:
	for _, q := range queries {
		idx, _ := slices.BinarySearch(obs, q[1])
		switch q[0] {
		case 1: // place obs
			obs = slices.Insert(obs, idx, q[1])
		case 2:
			maxGap := 0
			prevObs := 0
			for j := range idx {
				maxGap = max(maxGap, obs[j]-prevObs)
				prevObs = obs[j]
				if maxGap >= q[2] {
					res = append(res, true)
					continue outter
				}
			}
			maxGap = max(maxGap, q[1]-prevObs)
			res = append(res, maxGap >= q[2])
		}
	}
	return res
}

func getResults(queries [][]int) []bool {
	obs := make([]int, 0, 1<<5)
	res := make([]bool, 0, len(queries))
	st := NewSegmentTreeWithMap[int](15 * 1e4)
	for _, q := range queries {
		idx, _ := slices.BinarySearch(obs, q[1])
		switch q[0] {
		case 1: // place obs
			obs = slices.Insert(obs, idx, q[1])
			prevObs := 0
			if idx > 0 {
				prevObs = obs[idx-1]
			}
			st.Update(obs[idx], obs[idx]-prevObs)
			if idx < len(obs)-1 {
				st.Update(obs[idx+1], obs[idx+1]-obs[idx])
			}
		case 2:
			maxGap := q[1]
			if idx > 0 {
				maxGap = max(q[1]-obs[idx-1], st.Query(0, obs[idx-1]))
			}
			res = append(res, maxGap >= q[2])
		}
	}
	return res
}

type segmentTreeWithMap[T ~int] struct {
	data map[int]T
	size int
}

type SegmentTree[T ~int] interface {
	Query(lo, hi int) T
	Update(pos int, val T)
}

func NewSegmentTreeWithMap[T ~int](size int) SegmentTree[T] {
	return &segmentTreeWithMap[T]{
		data: make(map[int]T),
		size: size,
	}
}

func (st *segmentTreeWithMap[T]) Query(lo, hi int) T {
	lo += st.size
	hi += st.size
	res := st.data[hi]
	for lo < hi {
		if lo&1 != 0 {
			res = max(res, st.data[lo])
			lo++
		}
		if hi&1 != 0 {
			hi--
			res = max(res, st.data[hi])
		}
		lo >>= 1
		hi >>= 1
	}
	return res
}

func (st *segmentTreeWithMap[T]) Update(pos int, val T) {
	pos += st.size
	st.data[pos] = val

	for pos > 1 {
		pos >>= 1
		st.data[pos] = max(st.data[2*pos], st.data[2*pos+1])
	}
}
