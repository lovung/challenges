package mar2024

import "maps"

// https://leetcode.com/problems/subarrays-with-k-different-integers/
//
//	for(int l,r=0; r<n; r++){
//	    do_something_by_adding(nums[r]);
//	    while (!check_condition(k)){
//	        do_something_by_removing(nums[l];)
//	        l++;
//	    }
//	    update_the_answer();
//	}
func subarraysWithAtMostKDistinct(nums []int, k int) int {
	cnt := make(map[int]int)
	diff, result := 0, 0
	for l, r := 0, 0; r < len(nums); r++ {
		cnt[nums[r]]++
		if cnt[nums[r]] == 1 {
			diff++
		}
		for diff > k {
			cnt[nums[l]]--
			if cnt[nums[l]] == 0 {
				diff--
			}
			l++
		}
		result += r - l + 1
	}
	return result
}

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) -
		subarraysWithAtMostKDistinct(nums, k-1)
}

func subarraysWithKDistinct2(nums []int, k int) int {
	cnt, res := make(map[int]int), 0
	for l, r := 0, 0; r < len(nums); r++ {
		cnt[nums[r]]++
		for ; l < r && len(cnt) > k; l++ {
			if cnt[nums[l]] == 1 {
				delete(cnt, nums[l])
			} else {
				cnt[nums[l]]--
			}
		}
		if len(cnt) != k {
			continue
		}
		clone := maps.Clone(cnt)
		for i := l; i <= r && len(clone) == k; i++ {
			res++
			if clone[nums[i]] == 1 {
				delete(clone, nums[i])
			} else {
				clone[nums[i]]--
			}
		}
	}
	return res
}
