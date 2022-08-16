package algorithm

func twoSum(numbers []int, target int) []int {
	memoryMap := make(map[int]int)
	for i := range numbers {
		if oldIdx, ok := memoryMap[target-numbers[i]]; ok {
			return []int{oldIdx + 1, i + 1}
		}
		memoryMap[numbers[i]] = i
	}
	return nil
}

func twoSumV2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return nil
}
