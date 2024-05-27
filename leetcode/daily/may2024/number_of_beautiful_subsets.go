package may2024

import "sort"

// https://leetcode.com/problems/the-number-of-beautiful-subsets/
func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	cnt := make(map[int][]int)
	for _, num := range nums {
		mod := num % k
		cnt[mod] = append(cnt[mod], num)
	}
	ret := 1
	for _, kv := range cnt {
		tmp := 1
		dict := make(map[int]int)
		for _, num := range kv {
			v := tmp - dict[num-k]
			tmp += v
			dict[num] += v
		}
		ret *= tmp
	}
	return ret - 1
}

func beautifulSubsets1(nums []int, k int) int {
	n := 1 << len(nums)
	res := 0
	for i := 1; i < n; i++ {
		res += countIfBeautifulSubsets(nums, i, k)
	}
	return res
}

func countIfBeautifulSubsets(nums []int, bitmask, k int) int {
	exists := make(map[int]bool)
	bit := 1 << (len(nums) - 1)
	for i := range nums {
		if bitmask&bit != 0 {
			if exists[nums[i]-k] || exists[nums[i]+k] {
				return 0
			}
			exists[nums[i]] = true
		}
		bit >>= 1
	}
	return 1
}
