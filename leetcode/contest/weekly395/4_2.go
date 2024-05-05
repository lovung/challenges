package weekly395

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	cnt := make(map[int]int)
	for i := range nums {
		cnt[nums[i]]++
	}
	l, r := 1, len(cnt)
	for l < r {
		m := (l + r) / 2
		if atmost(nums, m)*2 >= total {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func atmost(nums []int, k int) int {
	res, i, n := 0, 0, len(nums)
	count := make(map[int]int)
	for j := range n {
		count[nums[j]]++
		for len(count) > k {
			count[nums[i]]--
			if count[nums[i]] == 0 {
				delete(count, nums[i])
			}
			i++
		}
		res += j - i + 1
	}
	return res
}
