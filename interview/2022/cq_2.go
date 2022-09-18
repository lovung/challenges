package interview

import "sort"

func findTotalImbalance(rank []int32) int64 {
	totalImbalance := int64(0)
	n := len(rank)
	for l := range rank {
		// reset the localImbalance and mark
		localImbalance := int64(0)
		// mark help remember the exist number in group,
		// we can use map here but the slice is more efficiency
		// because the roll number is unique and from 1 to n
		mark := make([]bool, n+2)
		mark[rank[l]] = true
		for r := l + 1; r < n; r++ {
			mark[rank[r]] = true
			switch {
			case mark[rank[r]+1] && mark[rank[r]-1]:
				localImbalance--
			case mark[rank[r]+1] || mark[rank[r]-1]:
				// don't need to recalculate the localImbalance
			default:
				localImbalance++
			}
			// Every group (i, j), we add the localImbalance to totalImbalance
			// instead of every i
			totalImbalance += localImbalance
		}
	}
	return totalImbalance
}

// brucefore -> timeout -> expected
func findTotalImbalance2(rank []int32) int64 {
	totalImbalance := int64(0)
	for i := range rank {
		for j := i + 2; j <= len(rank); j++ {
			totalImbalance += sortAndCountImbalance(rank[i:j])
		}
	}
	return totalImbalance
}

func sortAndCountImbalance(nums []int32) int64 {
	newArray := make([]int32, len(nums))
	copy(newArray, nums)
	sort.Slice(newArray, func(i, j int) bool {
		return newArray[i] < newArray[j]
	})
	imbalance := int64(0)
	for i := 1; i < len(newArray); i++ {
		if newArray[i]-newArray[i-1] > 1 {
			imbalance++
		}
	}
	return imbalance
}
