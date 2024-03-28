package problems

func canJump(nums []int) bool {
	n := len(nums)
	maxReachIdx := nums[0]
	for i := 1; i <= maxReachIdx && i < n; i++ {
		maxReachIdx = max(maxReachIdx, i+nums[i])
	}
	return maxReachIdx >= n-1
}

func jumpGame2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	maxReachIdx := nums[0]
	nextMaxIdx := maxReachIdx
	jumpCount := 1
	for i := 1; i <= maxReachIdx && i < n-1; i++ {
		if i <= nextMaxIdx {
			if maxReachIdx < i+nums[i] {
				maxReachIdx = i + nums[i]
			}
		}
		if i == nextMaxIdx {
			// fmt.Println(i, jumpCount, maxReachIdx)
			jumpCount++
			nextMaxIdx = maxReachIdx
		}
	}
	return jumpCount
}

// Jump Game 3
func canReach(arr []int, start int) bool {
	n := len(arr)
	arriveCount := make([]int, n)
	stack := make([]int, 0)

	for i := start; arr[i] != 0; {
		arriveCount[i]++
		if arriveCount[i] >= 2 {
			if len(stack) > 0 {
				// go to the index from
				i = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		// try to go to the new place with Least Freq Arrive
		if i+arr[i] < n && i-arr[i] >= 0 {
			if arriveCount[i+arr[i]] < arriveCount[i-arr[i]] {
				stack = append(stack, i-arr[i])
				i = i + arr[i]
			} else {
				stack = append(stack, i+arr[i])
				i = i - arr[i]
			}
		} else if i+arr[i] >= n && i-arr[i] >= 0 {
			i = i - arr[i]
		} else if i+arr[i] < n && i-arr[i] < 0 {
			i = i + arr[i]
		}
	}
	return true
}
