package weekly396

import "sort"

const mod = 1e9 + 7

func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
	sort.Ints(nums)
	if cost1*2 <= cost2 {
		return calculateSumGap(nums) * cost1 % mod
	}
	return minCostToEqualizeArray2(nums, cost1, cost2)
}

func calculateSumGap(nums []int) int {
	sumGap := 0
	maxVal := nums[len(nums)-1]
	for i := range nums {
		sumGap += maxVal - nums[i]
	}
	return sumGap
}

// When need to consider using cost2
func minCostToEqualizeArray2(nums []int, cost1, cost2 int) int {
	maxGap := nums[len(nums)-1] - nums[0]
	sumGap := calculateSumGap(nums)
	minCost := sumGap * cost1
	for k := 0; ; k++ {
		newCost := calculateCost(sumGap+k*len(nums), maxGap+k, cost1, cost2)
		if newCost <= minCost {
			minCost = newCost
		} else {
			break
		}
	}
	return minCost % mod
}

func calculateCost(sumGap, maxGap, cost1, cost2 int) int {
	maxCost2Time := sumGap - maxGap
	cost2Time := min(maxCost2Time, sumGap/2)
	cost1Time := sumGap - cost2Time*2
	cost := cost1Time*cost1 + cost2Time*cost2
	return cost
}
