package weekly395

import "slices"

func medianOfUniquenessArray1(nums []int) int {
	last := make([]int, 1e5+1)
	for i := range last {
		last[i] = -1
	}
	d := 0
	s := medianStore{}
	for i, n := range nums {
		if last[n] < 0 {
			d++
			s.add(d, i+1)
		} else {
			s.add(d, last[n]+1)
			s.add(d-1, i-last[n])
		}
		last[n] = i
	}
	return s.get()
}

type store struct {
	val, cnt int
}

type medianStore struct {
	data  []store
	total int
}

func (s *medianStore) add(val, cnt int) {
	idx, found := slices.BinarySearchFunc(s.data, val, func(_s store, t int) int {
		if _s.val == t {
			return 0
		}
		if _s.val > t {
			return -1
		}
		return 1
	})
	if found {
		s.data[idx].cnt += cnt
	} else {
		s.data = slices.Insert(s.data, idx, store{val, cnt})
		s.total++
	}
}

func (s *medianStore) get() int {
	cnt := s.total
	i := 0
	for cnt > 0 {
		cnt -= s.data[i].cnt
		i++
	}
	return s.data[i-1].val
}
