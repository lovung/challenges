package jul2024

import (
	"sort"
)

// https://leetcode.com/problems/sort-the-jumbled-numbers/
func sortJumbled(mapping []int, nums []int) []int {
	type node struct{ num, mapped int }
	nodes := make([]node, 0, len(nums))
	for _, num := range nums {
		nodes = append(nodes, node{num, convert(mapping, num)})
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		if nodes[i].mapped == nodes[j].mapped {
			return i < j
		}
		return nodes[i].mapped < nodes[j].mapped
	})
	res := make([]int, 0, len(nums))
	for _, n := range nodes {
		res = append(res, n.num)
	}
	return res
}

func convert(mapping []int, num int) int {
	if num == 0 {
		return mapping[0]
	}
	ans := 0
	e := 1
	for num > 0 {
		ans += (mapping[num%10]) * e
		e *= 10
		num /= 10
	}
	return ans
}
