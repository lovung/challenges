package problems

import "sort"

// PerPersonCost contains the cost and diff
type PersonCost struct {
	original []int
	diff     int
}

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	allPeople := make([]PersonCost, n)
	for i, e := range costs {
		allPeople[i].original = e
		allPeople[i].diff = e[1] - e[0]
	}

	quickSort(allPeople, 0, n-1)

	totalCost := 0
	for i, e := range allPeople {
		if i < n/2 {
			totalCost += e.original[1]
		} else {
			totalCost += e.original[0]
		}
	}
	return totalCost
}

func quickSort(arr []PersonCost, l, r int) {
	if l < r {
		pi := partition(arr, l, r)

		quickSort(arr, l, pi-1)
		quickSort(arr, pi+1, r)
	}
}

func partition(arr []PersonCost, l, r int) int {
	pivot := arr[r].diff

	i := l - 1
	for j := l; j <= r; j++ {
		if arr[j].diff < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

func twoCitySchedCostV2(costs [][]int) int {
	var (
		n  = len(costs)
		n2 = n / 2
	)
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][1]-costs[i][0] < costs[j][1]-costs[j][0]
	})
	totalCost := 0
	for i, e := range costs {
		if i < n2 {
			totalCost += e[1]
		} else {
			totalCost += e[0]
		}
	}

	return totalCost
}
