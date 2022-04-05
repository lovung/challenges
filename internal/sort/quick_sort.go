package sort

// QuickSort a Slice of interger
func QuickSort(A []int) {
	quickSort(A, 0, len(A)-1)
}

func quickSort(A []int, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := partition(A, lo, hi)
	quickSort(A, lo, p-1)
	quickSort(A, p+1, hi)
}

func partition(A []int, lo, hi int) int {
	pivot := A[hi]
	i := lo - 1
	for j := lo; j <= hi-1; j++ {
		if A[j] <= pivot {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	i++
	A[i], A[hi] = A[hi], A[i]
	return i
}
