package contest

func goodIndices(nums []int, k int) []int {
	leftLastOK := false
	rightLastOK := false
	n := len(nums)
	res := make([]int, 0, n)
	for i := k; i < n-k; i++ {
		if k == 1 {
			res = append(res, i)
			continue
		}
		curOK := 0
		if leftLastOK {
			if nums[i-1] <= nums[i-2] {
				curOK++
			} else {
				leftLastOK = false
			}
		} else {
			if checkNonIncreasing(nums[i-k : i]) {
				leftLastOK = true
				curOK++
			}
		}
		if rightLastOK {
			if nums[i+k] >= nums[i+k-1] {
				curOK++
			} else {
				rightLastOK = false
			}
		} else {
			if checkNonDecreasing(nums[i+1 : i+1+k]) {
				rightLastOK = true
				curOK++
			}
		}
		if curOK == 2 {
			res = append(res, i)
		}
	}
	return res
}

func checkNonIncreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			return false
		}
	}
	return true
}

func checkNonDecreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func goodIndices2(nums []int, k int) []int {
	n := len(nums)
	decreasing, increasing := make([]int, n), make([]int, n)
	decreasing[0] = 1
	increasing[n-1] = 1
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			decreasing[i] = 1 + decreasing[i-1]
		} else {
			decreasing[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			increasing[i] = 1 + increasing[i+1]
		} else {
			increasing[i] = 1
		}
	}
	res := make([]int, 0)
	for i := k; i < n-k; i++ {
		if decreasing[i-1] >= k && increasing[i+1] >= k {
			res = append(res, i)
		}
	}
	return res
}
