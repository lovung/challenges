package contest392

import (
	"strings"

	"golang.org/x/exp/constraints"
)

func getSmallestString(s string, k int) string {
	res := strings.Builder{}
	for i := 0; i < len(s); i++ {
		res.WriteByte(findSmallestCharWithinK(s[i], &k))
	}
	return res.String()
}

func findSmallestCharWithinK(c byte, k *int) byte {
	for i := byte('a'); i < 'z'; i++ {
		d := distance(c, i)
		if d <= *k {
			*k -= d
			return i
		}
	}
	return c
}

func distance(a, b byte) int {
	minD := 26
	minD = min(minD, abs(a, b))
	minD = min(minD, abs(a+26, b))
	minD = min(minD, abs(a, b+26))
	return minD
}

func abs[T constraints.Integer | constraints.Float | ~byte](a, b T) int {
	if a < b {
		return int(b - a)
	}
	return int(a - b)
}
