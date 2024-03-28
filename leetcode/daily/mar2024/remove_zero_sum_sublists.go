package mar2024

import (
	"slices"

	"github.com/lovung/ds/lists"
)

// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/description/
func removeZeroSumSublists(head *lists.ListNode[int]) *lists.ListNode[int] {
	arr := make([]int, 0)
	n := head
	for n != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}
	prefixSum := make([]int, len(arr)+1)
	prefixSum[0] = 0
	sum := 0
	for i := range arr {
		sum += arr[i]
		prefixSum[i+1] = sum
	}
	for {
		i, j, ok := findFirstLastEqual(prefixSum)
		if !ok {
			break
		}
		prefixSum = slices.Delete(prefixSum, i+1, j+1)
		arr = slices.Delete(arr, i, j)
	}
	return lists.NewLinkedListFromValues[int](arr)
}

func findFirstLastEqual(slice []int) (i, j int, ok bool) {
	exist := make(map[int]int)
	for j := range slice {
		if i, ok := exist[slice[j]]; ok {
			return i, j, true
		}
		exist[slice[j]] = j
	}
	return -1, -1, false
}

func removeZeroSumSublists2(head *lists.ListNode[int]) *lists.ListNode[int] {
	sum := 0
	prefixMap := make(map[int][]*lists.ListNode[int])
	for n := head; n != nil; n = n.Next {
		sum += n.Val
		if sum == 0 {
			head = n.Next
			continue
		}
		if lasts, ok := prefixMap[sum]; ok {
			for _, l := range lasts {
				l.Next = n.Next
			}
		}
		prefixMap[sum] = append(prefixMap[sum], n)
	}
	return head
}

func removeZeroSumSublists3(head *lists.ListNode[int]) *lists.ListNode[int] {
	prefixMap := make(map[int]*lists.ListNode[int])
	for n, sum := head, 0; n != nil; n = n.Next {
		sum += n.Val
		prefixMap[sum] = n
	}
	for n, sum := head, 0; n != nil; {
		sum += n.Val
		if sum == 0 {
			head, n = n.Next, n.Next
			continue
		}
		n.Next, n = prefixMap[sum].Next, n.Next
	}
	return head
}
