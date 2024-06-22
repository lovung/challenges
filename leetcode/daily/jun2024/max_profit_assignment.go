package jun2024

import (
	"slices"

	"github.com/lovung/ds/trees"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type job struct{ d, p int }
	jobs := make([]job, 0, len(difficulty))
	for i := range difficulty {
		jobs = append(jobs, job{difficulty[i], profit[i]})
	}
	slices.SortFunc(jobs, func(a, b job) int { return a.d - b.d })
	st := trees.NewSegmentTreeWithArray(len(jobs), func(a, b int) int { return max(a, b) })
	for i := range jobs {
		st.Update(i, jobs[i].p)
	}
	res := 0
	for _, w := range worker {
		idx, ok := slices.BinarySearchFunc(jobs, w, func(j job, ability int) int {
			return j.d - ability
		})
		for ok && idx < len(jobs) && jobs[idx].d == w {
			idx++
		}
		res += st.Query(0, idx)
	}
	return res
}
