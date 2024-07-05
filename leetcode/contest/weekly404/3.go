package weekly404

func maximumLength(nums []int, k int) int {
	mod := make([]int, len(nums))
	for i := range nums {
		mod[i] = nums[i] % k
	}
	firstIdx := make([]int, k)
	for i := range firstIdx {
		firstIdx[i] = -1
	}
	for i := range mod {
		if firstIdx[mod[i]] == -1 {
			firstIdx[mod[i]] = i
		}
	}
	nextMod := make([]int, k)
	for i := range nextMod {
		nextMod[i] = -1
	}
	nextModDp := make([][]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		nextModDp[i] = make([]int, k)
		copy(nextModDp[i], nextMod)
		nextMod[mod[i]] = i
	}

	res := 2
	for i := range k {
		for _, idx := range firstIdx {
			if idx == -1 {
				continue
			}
			cur := mod[idx]
			cnt := 1
			nextIdx := idx
			for nextIdx >= 0 {
				j := (i + k - cur) % k
				nextIdx = nextModDp[nextIdx][j]
				if nextIdx < 0 {
					break
				}
				cur = mod[nextIdx]
				cnt++
			}
			res = max(res, cnt)
		}
	}
	return res
}
