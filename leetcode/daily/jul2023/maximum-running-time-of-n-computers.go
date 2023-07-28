package jul2023

import (
	"container/heap"
	"sort"

	"github.com/lovung/ds/heaps"
)

// https://leetcode.com/problems/maximum-running-time-of-n-computers/
// Solution 1: Greedy
func maxRunTime(n int, batteries []int) int64 {
	maxHeap := heaps.MaxHeap[int]{}
	heap.Init(&maxHeap)

	res := int64(0)
	for i := range batteries {
		heap.Push(&maxHeap, batteries[i])
	}

	runningComputers := make([]int, n)
	for {
		for i := range runningComputers {
			runningComputers[i] = heap.Pop(&maxHeap).(int)
		}
		if runningComputers[len(runningComputers)-1] == 0 {
			return res
		}
		// smallest battery of runningComputers
		res++
		for i := range runningComputers {
			heap.Push(&maxHeap, runningComputers[i]-1)
		}
	}
}

// Solution 2: Sort batteries
func maxRunTime2(n int, batteries []int) int64 {
	length := len(batteries)
	sortedBatteries := make([]int, length)
	copy(sortedBatteries, batteries)
	sort.Ints(sortedBatteries)

	extraBat := 0
	for i := 0; i < length-n; i++ {
		extraBat += sortedBatteries[i]
	}

	startIdx := length - n
	for i := startIdx; i < length-1; i++ {
		needFillComputers := i - startIdx + 1
		gap := sortedBatteries[i+1] - sortedBatteries[i]
		if extraBat <= needFillComputers*gap {
			return int64(sortedBatteries[i] + extraBat/needFillComputers)
		}
		extraBat -= needFillComputers * gap
	}
	return int64(sortedBatteries[length-1] + extraBat/n)
}
