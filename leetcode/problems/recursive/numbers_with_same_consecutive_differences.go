package recursive

// Link: https://leetcode.com/problems/numbers-with-same-consecutive-differences/
func numsSameConsecDiff(n int, k int) []int {
	res := make([]int, 0)
	for i := 1; i <= 9; i++ {
		subProblem(n-1, &k, i, i, &res)
	}
	return res
}

func subProblem(lenCnt int, k *int, lastChar, curNumber int, res *[]int) {
	for i := 0; i <= 9; i++ {
		if lastChar-*k == i || lastChar+*k == i {
			if lenCnt == 1 {
				// build the full number
				*res = append(*res, curNumber*10+i)
			} else {
				// build the next charactor number
				subProblem(lenCnt-1, k, i, curNumber*10+i, res)
			}
		}
	}
}
