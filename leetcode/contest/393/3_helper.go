package contest393

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// find Least Common Multiple (LCM) via GCD
func LCM[T ~int | ~int64](integers ...T) T {
	if len(integers) == 1 {
		return integers[0]
	}
	a := integers[0]
	b := integers[1]
	result := a * b / gcd(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func SortUniq[T constraints.Ordered](s []T) []T {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	l := 0
	for r := 0; r < len(s); {
		for r < len(s) && s[l] == s[r] {
			r++
		}
		l++
		if r < len(s) {
			s[l] = s[r]
		}
	}
	return s[:l]
}
