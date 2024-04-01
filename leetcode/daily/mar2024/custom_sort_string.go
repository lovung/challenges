package mar2024

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/custom-sort-string
func customSortString(order string, s string) string {
	const a = 'a'
	priority := make([]int, 26)
	for i := range priority {
		priority[i] = 100
	}
	for i, o := range order {
		priority[o-a] = i
	}
	bs := []byte(s)
	sort.SliceStable(bs, func(i, j int) bool {
		return priority[bs[i]-a] < priority[bs[j]-a]
	})
	return string(bs)
}

func customSortString2(order string, s string) string {
	const a = 'a'
	counter := make([]int, 26)
	for i := range s {
		counter[s[i]-a]++
	}
	res := strings.Builder{}
	for i := range order {
		for j := 0; j < counter[order[i]-a]; j++ {
			res.WriteByte(order[i])
		}
		counter[order[i]-a] = 0
	}
	for i, c := range counter {
		for j := 0; j < c; j++ {
			res.WriteByte(byte(i) + a)
		}
	}
	return res.String()
}
