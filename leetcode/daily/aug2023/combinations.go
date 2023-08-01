package aug2023

// https://leetcode.com/problems/combinations/
func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	allItems := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		allItems = append(allItems, i)
	}
	recursiveCombine(allItems, []int{}, k, &ret)
	return ret
}

func recursiveCombine(remainItems []int, currentItems []int, k int, ret *[][]int) {
	if len(currentItems) == k {
		*ret = append(*ret, currentItems)
		return
	}
	if len(remainItems)+len(currentItems) < k {
		return
	}
	for i := range remainItems {
		newRemainItems := make([]int, len(remainItems)-i-1)
		copy(newRemainItems, remainItems[i+1:])

		newCurrentItems := make([]int, len(currentItems)+1)
		copy(newCurrentItems, currentItems)
		newCurrentItems[len(newCurrentItems)-1] = remainItems[i]
		recursiveCombine(newRemainItems, newCurrentItems, k, ret)
	}
}
