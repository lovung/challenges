// Given an array of integers of size >= 2, find its two largest elements

// int64
// 3 elements [0, 1, 1, 1] -> [1, 1]
//

package coupang

// Output is DESC (first -> second)
func findTwoLargest(input []int64) []int64 {
	// Solution 1: heap O(n*log(n))
	// Solution 2: 2 variables O(n)
	// - New item which >= first: first <- item; second <- first
	// - Else new item which > second: second <- item
	first, second := input[0], input[1]
	if input[0] < input[1] {
		first, second = second, first
	}

	for i := 2; i < len(input); i++ {
		if input[i] >= first {
			first, second = input[i], first
		} else if input[i] > second {
			second = input[i]
		}
	}
	return []int64{first, second}
}
