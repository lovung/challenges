package contest

const mod = 1e9 + 7

type cachePair struct {
	pos, k int
}

// Link: https://leetcode.com/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps
func numberOfWays(startPos int, endPos int, k int) int {
	cache := make(map[cachePair]int)
	return recursive(startPos, &endPos, k, cache) % mod
}

func recursive(curPos int, endPos *int, curK int, cache map[cachePair]int) int {
	if val, ok := cache[cachePair{curPos, curK}]; ok {
		return val
	}
	if curK == 1 {
		if curPos-*endPos == 1 || curPos-*endPos == -1 {
			return 1
		}
		return 0
	}
	if curPos == *endPos && curK%2 != 0 {
		return 0
	}
	if curPos < *endPos {
		if *endPos-curPos > curK {
			return 0
		} else if *endPos-curPos == curK {
			return 1
		}
	} else {
		if curPos-*endPos > curK {
			return 0
		} else if curPos-*endPos == curK {
			return 1
		}
	}

	res := recursive(curPos+1, endPos, curK-1, cache) + recursive(curPos-1, endPos, curK-1, cache)%mod
	cache[cachePair{curPos, curK}] = res
	return res
}
