package mar2024

// https://leetcode.com/problems/count-elements-with-maximum-frequency
func maxFrequencyElements(nums []int) int {
	freqMatrix := make([][]int, len(nums)+1)
	for i := range nums {
		meet(nums[i], freqMatrix)
	}
	for i := len(freqMatrix) - 1; i > 0; i-- {
		if len(freqMatrix[i]) > 0 {
			return len(freqMatrix[i]) * i
		}
	}
	return 0
}

func meet(n int, freqMatrix [][]int) {
	for i := 1; i < len(freqMatrix)-1; i++ {
		for j := range freqMatrix[i] {
			if freqMatrix[i][j] == n {
				copy(freqMatrix[i][j:], freqMatrix[i][j+1:])
				freqMatrix[i] = freqMatrix[i][:len(freqMatrix[i])-1]
				freqMatrix[i+1] = append(freqMatrix[i+1], n)
				return
			}
		}
	}
	freqMatrix[1] = append(freqMatrix[1], n)
}

func maxFrequencyElements2(nums []int) int {
	freq := make([]int, 101)
	maxFreq, res := 0, 0
	for _, n := range nums {
		freq[n]++
		if freq[n] > maxFreq {
			maxFreq = freq[n]
			res = freq[n]
		} else if freq[n] == maxFreq {
			res += freq[n]
		}
	}
	return res
}
