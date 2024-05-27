package weekly398

// Timeout
func waysToReachStair(k int) int {
	cnt := 0
	recursive(0, 1, k, 0, &cnt)
	return cnt
}

func recursive(prevPos, curPos, target, jump int, cnt *int) {
	if curPos > target+1 {
		return
	}
	if curPos == target {
		*cnt++
	}
	if prevPos < curPos {
		recursive(curPos, curPos-1, target, jump, cnt)
	}
	recursive(curPos, curPos+pow2(jump), target, jump+1, cnt)
}

func pow2(n int) int {
	res := 1
	for range n {
		res <<= 1
	}
	return res
}

func waysToReachStair2(k int) int {
	res := 0
	for j := range 31 {
		res += comb(j+1, (1<<j)-k)
	}
	return res
}

func comb(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	res := 1
	for i := range k {
		res = res * (n - i) / (i + 1)
	}
	return res
}
