package mar2024

import "math"

// https://leetcode.com/problems/task-scheduler
func leastInterval(tasks []byte, n int) int {
	const a = 'A'
	counter := make([]int, 26)
	for i := range tasks {
		counter[tasks[i]-a]++
	}
	lastIdx := make([]int, 26)
	for i := range lastIdx {
		lastIdx[i] = math.MinInt
	}
	curStep := 0
	for sum(counter) > 0 {
		task, maxCnt := chooseTask(lastIdx, counter, curStep-n)
		if maxCnt > 0 {
			// Enough idle
			lastIdx[task] = curStep
			counter[task]--
		}
		curStep++
	}
	return curStep
}

func sum(s []int) int {
	ret := 0
	for i := range s {
		ret += s[i]
	}
	return ret
}

func chooseTask(lastIdx, counter []int, target int) (int, int) {
	idx, maxCnt := -1, 0
	for i := range counter {
		if counter[i] > maxCnt && lastIdx[i] < target {
			maxCnt = counter[i]
			idx = i
		}
	}
	return idx, maxCnt
}

func leastInterval2(tasks []byte, n int) int {
	count := make([]int, 26)
	for i := range tasks {
		count[tasks[i]-'A']++
	}

	maxCount := count[0]
	for j := range count {
		maxCount = max(maxCount, count[j])
	}
	i := 0
	for j := range count {
		if count[j] == maxCount {
			i++
		}
	}
	return max(len(tasks), (maxCount-1)*(n+1)+i)
}
