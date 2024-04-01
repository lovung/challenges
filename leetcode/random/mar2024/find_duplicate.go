package mar2024

// https://leetcode.com/problems/find-the-duplicate-number/
func findDuplicate_hashmap(nums []int) int {
	exist := make([]bool, len(nums)+1)
	for i := range nums {
		if exist[nums[i]] {
			return nums[i]
		}
		exist[nums[i]] = true
	}
	return 1
}

func findDuplicate_bubbleSort(nums []int) int {
	n := make([]int, len(nums))
	copy(n, nums)
	for i := 0; i < len(n); {
		if n[i] == i+1 {
			i++
			continue
		}
		if n[i] == n[n[i]-1] {
			return n[i]
		}
		n[i], n[n[i]-1] = n[n[i]-1], n[i]
	}
	return 1
}

func findDuplicate_tortoiseAndHare(nums []int) int {
	s, f := nums[0], nums[0]
	for {
		s, f = nums[s], nums[nums[f]]
		if s == f {
			break
		}
	}
	s = nums[0]
	for s != f {
		s, f = nums[s], nums[f]
	}
	return f
}
