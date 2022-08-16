package june2022

func singleNumber(nums []int) int {
	mm := make(map[int]int)
	for i := range nums {
		if _, ok := mm[nums[i]]; ok {
			delete(mm, nums[i])
		} else {
			mm[nums[i]] = 1
		}
	}
	for k := range mm {
		return k
	}
	return 0
}

func singleNumber2(nums []int) int {
	mm := make(map[int]int)
	for i := range nums {
		if _, ok := mm[nums[i]]; ok {
			mm[nums[i]]++
		} else {
			mm[nums[i]] = 1
		}
		if mm[nums[i]] == 3 {
			delete(mm, nums[i])
		}
	}
	for k := range mm {
		return k
	}
	return 0
}

func singleNumber3(nums []int) []int {
	mm := make(map[int]int)
	for i := range nums {
		if _, ok := mm[nums[i]]; ok {
			delete(mm, nums[i])
		} else {
			mm[nums[i]] = 1
		}
	}
	ret := make([]int, 0, 2)
	for k := range mm {
		ret = append(ret, k)
	}
	return ret
}
