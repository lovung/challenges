package jun2024

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2MapIdx := make(map[int]int)
	for i, v := range arr2 {
		arr2MapIdx[v] = i
	}
	notExist := make([]int, 0, len(arr1))
	exist := make([]int, 0, len(arr1))
	for _, v := range arr1 {
		if _, ok := arr2MapIdx[v]; !ok {
			notExist = append(notExist, v)
		} else {
			exist = append(exist, v)
		}
	}
	sort.SliceStable(exist, func(i, j int) bool {
		return arr2MapIdx[exist[i]] < arr2MapIdx[exist[j]]
	})
	sort.Ints(notExist)
	return append(exist, notExist...)
}

func relativeSortArray2(arr1 []int, arr2 []int) []int {
	arr2MapIdx := make(map[int]int)
	for i, v := range arr2 {
		arr2MapIdx[v] = i
	}
	r := len(arr1) - 1
	for l := 0; l <= r; {
		if _, ok := arr2MapIdx[arr1[l]]; !ok {
			arr1[l], arr1[r] = arr1[r], arr1[l]
			r--
		} else {
			l++
		}
	}
	sort.SliceStable(arr1[:r+1], func(i, j int) bool {
		return arr2MapIdx[arr1[i]] < arr2MapIdx[arr1[j]]
	})
	sort.Ints(arr1[r+1:])
	return arr1
}
