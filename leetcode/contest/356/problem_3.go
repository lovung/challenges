package contest356

import "strings"

func minimumString(a string, b string, c string) string {
	allMerged := [6]string{
		tryMerge(tryMerge(a, b), c),
		tryMerge(tryMerge(a, c), b),
		tryMerge(tryMerge(b, a), c),
		tryMerge(tryMerge(b, c), a),
		tryMerge(tryMerge(c, b), a),
		tryMerge(tryMerge(c, a), b),
	}

	ret := allMerged[0]
	for i := 1; i < 6; i++ {
		if len(allMerged[i]) < len(ret) ||
			(len(allMerged[i]) == len(ret) && allMerged[i] < ret) {
			ret = allMerged[i]
		}
	}
	return ret
}

func tryMerge(first, second string) string {
	if strings.Contains(first, second) {
		return first
	} else if strings.Contains(second, first) {
		return second
	}
	merged := cntMergeLen(first, second)
	ret := first + second[merged:]
	return ret
}

func cntMergeLen(first, second string) int {
	lessLen := len(first)
	if lessLen > len(second) {
		lessLen = len(second)
	}
	max := 0
	for i := 0; i <= lessLen; i++ {
		a := first[len(first)-i:]
		b := second[:i]
		if a == b {
			max = i
		}
	}
	return max
}
