package oct2022

import (
	"container/heap"
	"sort"

	"github.com/lovung/ds/heaps"
)

type task struct {
	name       string
	start, run int
	priority   int
}

func executeTask(tasks []*task) []string {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].start < tasks[j].start
	})

	curIndex := 0
	pivot := 0
	res := make([]string, 0, len(tasks))
	maxHeap := heaps.MaxHeapWithValue[int]{}
	heap.Init(&maxHeap)

	for len(res) < len(tasks) {
		for curIndex < len(tasks) {
			if tasks[curIndex].start <= pivot {
				heap.Push(&maxHeap, &heaps.HeapItem[int]{
					Ref:   tasks[curIndex].priority,
					Value: tasks[curIndex],
				})
				curIndex++
			} else {
				break
			}
		}

		if maxHeap.Len() == 0 {
			pivot = tasks[curIndex+1].start
			continue
		}

		item := heap.Pop(&maxHeap).(*heaps.HeapItem[int])
		highestPriTask := item.Value.(*task)
		res = append(res, highestPriTask.name)
		pivot += highestPriTask.run
	}
	return res
}
