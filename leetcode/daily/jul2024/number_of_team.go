package jul2024

// https://leetcode.com/problems/count-number-of-teams/description/
func numTeams(rating []int) int {
	increase := numTeamsIncrease(rating)
	revert(rating)
	decrease := numTeamsIncrease(rating)
	return increase + decrease
}

func revert[T any](slice []T) {
	n := len(slice)
	for i := range n >> 1 {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
}

func numTeamsIncrease(rating []int) int {
	const N = 1e5 + 1
	mark := make([]bool, N+1)
	cnt := make([]int, N+1)
	for i := len(rating) - 1; i >= 0; i-- {
		rate := rating[i]
		for r := range rate {
			if !mark[r] {
				cnt[r]++
			}
		}
		mark[rate] = true
	}
	res := 0
	for i, r := range rating {
		if cnt[r] < 2 {
			continue
		}
		for j := i + 1; j < len(rating)-1; j++ {
			if rating[j] < rating[i] {
				continue
			}
			res += cnt[rating[j]]
		}
	}
	return res
}
