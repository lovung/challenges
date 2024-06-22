package jun2024

// https://leetcode.com/problems/grumpy-bookstore-owner
// Sliding window
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	satisfiedCnt := 0
	for i := range customers {
		if grumpy[i] == 0 {
			satisfiedCnt += customers[i]
		}
	}
	l, r := 0, 0
	unsatisfiedCnt := 0
	for ; r < minutes; r++ {
		if grumpy[r] == 1 {
			unsatisfiedCnt += customers[r]
		}
	}
	maxUnsa := unsatisfiedCnt
	for r < len(customers) {
		if grumpy[r] == 1 {
			unsatisfiedCnt += customers[r]
		}
		if grumpy[l] == 1 {
			unsatisfiedCnt -= customers[l]
		}
		maxUnsa = max(maxUnsa, unsatisfiedCnt)
		r++
		l++
	}
	return satisfiedCnt + maxUnsa
}
