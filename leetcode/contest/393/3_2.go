package contest393

func findKthSmallest_2(coins []int, k int) int64 {
	lcm := int64(LCM(coins...))
	res := int64(0)
	arr := make([]int64, 0, 100)
	for i := int64(1); i <= lcm; i++ {
		if isOneMultiple(i, coins) {
			res = i
			arr = append(arr, i)
		}
	}
	milestone := int64((k-1)/len(arr)) * res
	index := (k - 1) % len(arr)
	return milestone + arr[index]
}

func isOneMultiple(n int64, coins []int) bool {
	for i := range coins {
		if n%int64(coins[i]) == 0 {
			return true
		}
	}
	return false
}
