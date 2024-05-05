package contest393

func maximumPrimeDifference(nums []int) int {
	minPrime, maxPrime := 101, 0
	for i, n := range nums {
		if isPrime(n) {
			minPrime = min(minPrime, i)
			maxPrime = max(maxPrime, i)
		}
	}
	return maxPrime - minPrime
}

func isPrime(n int) bool {
	for i := range primes {
		if n == primes[i] {
			return true
		}
	}
	return false
}

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
