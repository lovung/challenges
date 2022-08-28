package contest

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	sumArray := make([]int, len(nums)+1)
	sum := 0
	for i := range nums {
		sum += nums[i]
		sumArray[i+1] = sum
	}
	fmt.Println(sumArray)
	ans := make([]int, 0, len(queries))
	for i := range queries {
		index := binarySearch(sumArray, func(j int) int {
			if queries[i] < sumArray[j] {
				return -1
			} else if queries[i] > sumArray[j] {
				return 1
			}
			return 0
		})
		ans = append(ans, index)
	}
	return ans
}

func binarySearch(slice []int, compFunc func(i int) int) int {
	l, r := 0, len(slice)-1
	pivot := 0
	for l <= r {
		pivot = (l + r) / 2
		cmpVal := compFunc(pivot)
		switch cmpVal {
		case -1:
			r = pivot - 1
		case 1:
			l = pivot + 1
		case 0:
			return pivot
		}
	}
	return l-1
}
