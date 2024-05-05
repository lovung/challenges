package contest

// https://leetcode.com/problems/max-pair-sum-in-an-array/
func maxSum(nums []int) int {
	mapMaxDigit := make(map[int]int)
	ret := -1
	for _, n := range nums {
		maxD := maxDigit(n)
		if mapMaxDigit[maxD] == 0 {
			mapMaxDigit[maxD] = n
			continue
		}
		ret = max(ret, mapMaxDigit[maxD]+n)
		mapMaxDigit[maxD] = max(mapMaxDigit[maxD], n)
	}
	return ret
}

func maxDigit(num int) int {
	maxD := 0
	for num > 0 {
		maxD = max(maxD, num%10)
		num /= 10
	}
	return maxD
}
