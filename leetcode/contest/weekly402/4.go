package weekly402

import "github.com/lovung/ds/trees"

func countOfPeaks(nums []int, queries [][]int) []int {
	st := trees.NewSegmentTreeWithArray(len(nums), func(a, b int) int {
		return a + b
	})
	for i := range nums {
		update(nums, i, st)
	}
	res := make([]int, 0, len(queries))
	for _, q := range queries {
		if q[0] == 1 {
			res = append(res,
				max(0,
					st.Query(q[1], q[2]+1)-
						st.Query(q[1], q[1]+1)-
						st.Query(q[2], q[2]+1)))
		} else {
			nums[q[1]] = q[2]
			for i := q[1] - 1; i <= q[1]+1; i++ {
				update(nums, i, st)
			}
		}
	}
	return res
}

func update(nums []int, i int, st trees.SegmentTree[int]) {
	if i <= 0 || i >= len(nums)-1 {
		return
	}
	if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
		st.Update(i, 1)
	} else {
		st.Update(i, 0)
	}
}
