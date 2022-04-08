package math

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a ...T) T {
	max := a[0]
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func Min[T constraints.Ordered](a ...T) T {
	min := a[0]
	for i := range a {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}
