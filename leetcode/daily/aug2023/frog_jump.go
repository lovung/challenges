package aug2023

import "github.com/lovung/ds/queue"

type bfsItem struct {
	index    int
	lastJump int
}

func canCross(stones []int) bool {
	bfs := queue.NewSimpleQueue[*bfsItem]()
	bfs.EnQueue(&bfsItem{1, stones[1] - stones[0]})
	cache := make(map[*bfsItem]bool)
	for bfs.Len() > 0 {
		item, _ := bfs.DeQueue()
		if item.index == len(stones)-1 {
			return true
		}
		if cache[item] == true {
			continue
		}
		for i := item.index + 1; i < len(stones); i++ {
			// because its next jump must be either k - 1, k, or k + 1 units
			if stones[i]-stones[item.index] > item.lastJump+1 {
				break
			}
			if canJump(stones, item.index, i, item.lastJump) {
				bfs.EnQueue(&bfsItem{i, stones[i] - stones[item.index]})
			}
		}
		cache[item] = true
	}
	return false
}

func canJump(stones []int, curIdx, nextIdx, lastJump int) bool {
	if nextIdx >= len(stones) {
		return false
	}
	switch stones[nextIdx] - stones[curIdx] - lastJump {
	case -1, 0, 1:
		return true
	default:
		return false
	}
}
