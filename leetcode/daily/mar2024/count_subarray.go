package mar2024

// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/
//
//	for(int l,r=0; r<n; r++){
//	    do_something_by_adding(nums[r]);
//	    while (!check_condition(k)){
//	        do_something_by_removing(nums[l];)
//	        l++;
//	    }
//	    update_the_answer();
//	}
func countSubarrays(nums []int, k int) int64 {
	maxElem := 0
	for i := range nums {
		maxElem = max(maxElem, nums[i])
	}
	cnt, n, ret := 0, len(nums), int64(0)
	for l, r := 0, 0; r < n; r++ {
		if nums[r] == maxElem {
			cnt++
		}
		for ; cnt >= k; l++ {
			if nums[l] == maxElem {
				cnt--
			}
		}
		ret += int64(l)
	}
	return ret
}

func countSubarrays2(nums []int, k int) int64 {
	maxElem := 0
	for i := range nums {
		maxElem = max(maxElem, nums[i])
	}
	ret := int64(0)
	queue := make([]int, 0, len(nums))
	for i := range nums {
		if nums[i] == maxElem {
			queue = append(queue, i)
		}

		if len(queue) >= k {
			ret += int64(queue[len(queue)-k] + 1)
		}
	}
	return ret
}

// 1 -> n
// 2 -> n-1
// ...
// n -> 1
// => 1+2+...+n
// => n*(n+1)/2

// 1,3,2,3,3; k=2
// [1,3,2,3], [3,2,3]
// [1,3,2,3,3], [3,2,3,3], [2,3,3] and [3,3]
