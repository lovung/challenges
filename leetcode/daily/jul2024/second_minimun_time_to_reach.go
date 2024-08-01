package jul2024

import (
	"container/heap"

	"github.com/emirpasic/gods/v2/queues/linkedlistqueue"
	"github.com/lovung/ds/heaps"
)

// https://leetcode.com/problems/second-minimum-time-to-reach-destination/
// Wrong answer
func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	q := heaps.MinHeapWithValue[int]{}
	heap.Push(&q, &heaps.HeapItem[int]{Ref: 0, Value: 1})
	var first bool
	for q.Len() > 0 {
		node := heap.Pop(&q).(*heaps.HeapItem[int])
		for _, near := range adj[node.Value.(int)] {
			if near == n {
				if !first {
					first = true
				} else {
					return node.Ref + time
				}
			}
			heap.Push(&q, &heaps.HeapItem[int]{
				Ref:   calcLeaveNodeTime(node.Ref, time, change),
				Value: near,
			})
		}
	}
	return -1
}

func calcLeaveNodeTime(prev, time, change int) int {
	arriveTime := prev + time
	if (arriveTime/change)%2 == 0 {
		return arriveTime
	}
	return (arriveTime/change)*change + change
}

func secondMinimum2(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	q := linkedlistqueue.New[int]()
	q.Enqueue(1)
	curTime := 0
	res := -1
	visitTimes := make([][]int, n+1)

	for !q.Empty() {
		for i := q.Size(); i > 0; i-- {
			node, _ := q.Dequeue()
			if node == n {
				if res != -1 {
					return curTime
				}
				res = curTime
			}

			for _, nei := range adj[node] {
				neiTimes := visitTimes[nei]
				if len(neiTimes) == 0 ||
					(len(neiTimes) == 1 && neiTimes[0] != curTime) {
					q.Enqueue(nei)
					visitTimes[nei] = append(visitTimes[nei], curTime)
				}
			}
		}

		if (curTime/change)%2 == 1 {
			curTime += change - (curTime % change)
		}
		curTime += time
	}

	return -1
}
