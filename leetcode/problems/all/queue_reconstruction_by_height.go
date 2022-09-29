package problems

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	var (
		n   = len(people)
		out = make([][]int, n)
	)
	for i := 0; i < n; i++ {
		out[i] = make([]int, 2)
		out[i][0] = -1
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0]
	})
	for _, e := range people {
		k := e[1]
		for j := 0; j < n; j++ {
			if out[j][0] != -1 && out[j][0] < e[0] {
				continue
			}

			if k == 0 {
				out[j][0] = e[0]
				out[j][1] = e[1]
				break
			}
			k--
		}
	}
	return out
}
