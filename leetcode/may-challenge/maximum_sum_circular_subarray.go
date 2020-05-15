package may-challenge

import "fmt"
import "math"

func maxSubarraySumCircular(A []int) int {
	maxSum := 0
	max := math.MinInt64
	minSum := 0
	min := math.MaxInt64
	total := 0
	fmt.Println(A)
	for _, e := range A {
		maxSum += e
		if maxSum < e {
			maxSum = e
		}

		if max < maxSum {
			max = maxSum
		}

		minSum += e
		if minSum > e {
			minSum = e
		}
		if min > minSum {
			min = minSum
		}
		total += e
		// fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", e, maxSum, max, minSum, min, total)
	}

	if max < 0 {
		return max
	}

	tmp := total - min
	if tmp > max {
		return tmp
	}
	return max
}

// func maxSubarraySumCircular(A []int) int {
// 	var sum, maxE int
// 	nA := append(A, A...)
// 	n := len(A)
// 	prefixSum := make([]int, 2*n+1)
// 	prefixSum[0] = 0
// 	maxE = math.MinInt64
// 	// minSum = math.MaxInt64

// 	for i := 0; i < 2*n; i++ {
// 		sum += nA[i]
// 		prefixSum[i+1] = sum
// 		if nA[i] > maxE {
// 			maxE = nA[i]
// 		}
// 	}
// 	sum >>= 1

// 	fmt.Println(prefixSum, sum, maxE)
// 	max := math.MinInt64
// 	maxIndex := 2*n + 1
// 	isNegative := true
// 	out := sum
// 	for maxIndex > n {
// 		for i := maxIndex - 1; i > 0; i-- {
// 			if maxIndex != -1 && maxIndex-i >= n {
// 				break
// 			}
// 			if prefixSum[i] > max {
// 				max = prefixSum[i]
// 				maxIndex = i
// 			}

// 			tmp := max - prefixSum[i]
// 			fmt.Println(max, maxIndex, tmp)
// 			if tmp > out {
// 				out = tmp
// 			}
// 			if tmp != 0 {
// 				isNegative = false
// 			}
// 		}
// 		maxIndex--
// 		max = prefixSum[maxIndex]
// 	}

// 	if isNegative {
// 		return maxE
// 	}
// 	return out
// }

// import "math"

// func maxSubarraySumCircular(A []int) int {
// 	length := len(A)
// 	prefixSum := make([]int, length+1)
// 	prefixSum[0] = 0
//     out := math.MinInt64
// 	var sum, r, newSum int

// 	for i := 0; i < length; i++ {
// 		sum += A[i]
// 		prefixSum[i+1] = sum
//         if A[i] > out {
//             out = A[i]
//         }
// 	}

// 	for k := 2; k <= length; k++ {
// 		for l := 0; l < length; l++ {
// 			r = l + k - 1
// 			if r < length {
// 				newSum = prefixSum[r+1] - prefixSum[l]
// 				if newSum > out {
// 					out = newSum
// 				}
// 			} else {
// 				r -= length
// 				newSum = sum - (prefixSum[l] - prefixSum[r+1])
// 				if newSum > out {
// 					out = newSum
// 				}
// 			}
// 		}
// 	}
// 	return out
// }
