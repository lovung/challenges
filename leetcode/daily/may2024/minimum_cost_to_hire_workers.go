package may2024

import (
	"container/heap"
	"math"
	"slices"
	"sort"
)

type fraction struct {
	q, w int
}

// https://leetcode.com/problems/minimum-cost-to-hire-k-workers
func mincostToHireWorkers1(quality []int, wage []int, k int) float64 {
	workers := make([]fraction, 0, len(quality))
	for i := range quality {
		workers = append(workers, fraction{quality[i], wage[i]})
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].q*workers[j].w < workers[i].w*workers[j].q
	})
	res := float64(math.MaxFloat64)
	for i := range len(quality) - k + 1 {
		paid := calculateMinPaid(workers, k, i)
		res = min(res, paid)
	}
	return res
}

func calculateMinPaid(worker []fraction, k, start int) float64 {
	newWorkers := make([]fraction, len(worker)-start)
	copy(newWorkers, worker[start:])
	needSortSlice := newWorkers[1:]
	sort.Slice(needSortSlice, func(i, j int) bool {
		return needSortSlice[i].q < needSortSlice[j].q
	})
	paid := float64(0)
	base := newWorkers[0]
	for j := range k {
		paid += float64(base.w) * float64(newWorkers[j].q) / float64(base.q)
	}
	return paid
}

func mincostToHireWorkers2(quality []int, wage []int, k int) float64 {
	n := len(quality)
	workers := make([]fraction, 0, n)
	for i := range quality {
		workers = append(workers, fraction{quality[i], wage[i]})
	}
	slices.SortFunc(workers, func(a, b fraction) int { return a.q*b.w - a.w*b.q })
	quality = slices.Clone(quality)
	res := float64(math.MaxFloat64)
	sort.Ints(quality)
	for i := range len(quality) - k + 1 {
		idx, _ := slices.BinarySearch(quality, workers[i].q)
		quality = slices.Delete(quality, idx, idx+1)
		paid := calculateMinPaid2(workers[i], quality[:k-1])
		res = min(res, paid)
	}
	return res
}

func calculateMinPaid2(base fraction, quality []int) float64 {
	paid := float64(base.w)
	for j := range quality {
		paid += float64(base.w) * float64(quality[j]) / float64(base.q)
	}
	return paid
}

func mincostToHireWorkers3(quality, wage []int, k int) float64 {
	workers := make([]fraction, len(quality))
	for i, q := range quality {
		workers[i] = fraction{q, wage[i]}
	}
	slices.SortFunc(workers, func(a, b fraction) int { return a.w*b.q - b.w*a.q })
	h := hp{make([]int, k)}
	sumQ := 0
	for i, p := range workers[:k] {
		h.IntSlice[i] = p.q
		sumQ += p.q
	}
	heap.Init(&h)
	ans := float64(sumQ*workers[k-1].w) / float64(workers[k-1].q)

	for _, p := range workers[k:] {
		if p.q < h.IntSlice[0] {
			sumQ -= h.IntSlice[0] - p.q
			h.IntSlice[0] = p.q
			heap.Fix(&h, 0)
			ans = min(ans, float64(sumQ*p.w)/float64(p.q))
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
