package jul2024

import "sort"

// https://leetcode.com/problems/sort-array-by-increasing-frequency/
func frequencySort(nums []int) []int {
	cnt := make(map[int]int)
	for _, n := range nums {
		cnt[n]++
	}
	type freq struct{ n, cnt int }
	freqs := make([]freq, 0, len(cnt))
	for k, n := range cnt {
		freqs = append(freqs, freq{k, n})
	}
	sort.Slice(freqs, func(i, j int) bool {
		if freqs[i].cnt == freqs[j].cnt {
			return freqs[i].n > freqs[j].n
		}
		return freqs[i].cnt < freqs[j].cnt
	})
	res := make([]int, 0, len(nums))
	for _, f := range freqs {
		for range f.cnt {
			res = append(res, f.n)
		}
	}
	return res
}
