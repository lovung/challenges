package medium

import (
	"container/heap"
	"sort"

	"github.com/lovung/challenges/internal/heaps"
)

// Link: https://leetcode.com/problems/next-permutation
/*
	A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

	For example, for arr = [1,2,3], the following are considered permutations of arr: [1,2,3], [1,3,2], [3,1,2], [2,3,1].
	The next permutation of an array of integers is the next lexicographically greater permutation of its integer. More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

	For example, the next permutation of arr = [1,2,3] is [1,3,2].
	Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
	While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
	Given an array of integers nums, find the next permutation of nums.

	[1, 2, 3] -> [1, 3, 2]
	[1, 3, 2] -> [2, 1, 3]
	[2, 1, 3] -> [2, 3, 1]
	[2, 3, 1] -> [3, 1, 2]
	[3, 1, 2] -> [3, 2, 1]
	[3, 2, 1] -> [1, 2, 3]

	[4, 1, 3, 2] -> [4, 2, 1, 3]
	[4, 0, 1, 3, 2] -> [4, 0, 2, 1, 3]


	[5,4,7,5,3,2]

	The replacement must be in place and use only constant extra memory.
*/
func nextPermutation(nums []int) {
	h := heaps.MinHeapWithValue[int]{}
	heap.Init(&h)
	i := len(nums) - 1
	for ; i > 0; i-- {
		item := &heaps.HeapItem[int]{
			Ref:   nums[i],
			Value: i,
		}
		heap.Push(&h, item)
		if nums[i-1] < nums[i] {
			minValue := heap.Pop(&h)
			for minValue != nil && minValue.(*heaps.HeapItem[int]).Ref <= nums[i-1] {
				minValue = heap.Pop(&h)
			}
			gotcha := minValue.(*heaps.HeapItem[int])
			nums[i-1], nums[gotcha.Value.(int)] = nums[gotcha.Value.(int)], nums[i-1]
			break
		}
	}
	sort.Ints(nums[i:])
}
