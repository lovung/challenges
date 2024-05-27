package weekly398

import "github.com/lovung/ds/trees"

func isArraySpecial(nums []int, queries [][]int) []bool {
	res := make([]bool, 0, len(queries))
	st := trees.NewSegmentTreeWithArray(len(nums), func(a, b bool) bool {
		return a && b
	})
	st.SetInitQueryValue(true)
	for i := range nums {
		if i == 0 {
			st.Update(i, true)
		} else {
			st.Update(i, nums[i]&1 != nums[i-1]&1)
		}
	}
	for _, q := range queries {
		i := st.Query(q[0]+1, q[1]+1)
		res = append(res, i)
	}
	return res
}
