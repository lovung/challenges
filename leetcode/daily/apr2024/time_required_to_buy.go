package apr2024

// https://leetcode.com/problems/time-needed-to-buy-tickets/description/
func timeRequiredToBuy(tickets []int, k int) int {
	cnt := 0
	for i := range tickets {
		if i <= k {
			cnt += min(tickets[k], tickets[i])
		} else {
			cnt += min(tickets[k]-1, tickets[i])
		}
	}
	return cnt
}
