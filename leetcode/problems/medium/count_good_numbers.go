package medium

const mod = 10e8 + 7

func countGoodNumbers(n int64) int {
	return int(splitCount(n) % mod)
}

func splitCount(n int64) int64 {
	switch n {
	case 1:
		return 5
	case 2:
		return 20
	case 3:
		return 100
	}
	switch n % 4 {
	case 0:
		tmp := splitCount(n / 2)
		return tmp * tmp % mod
	case 1:
		tmp := splitCount(n / 2)
		return tmp * 5 % mod * tmp % mod
	case 2:
		tmp := splitCount(n/2 - 1)
		return tmp * 20 % mod * tmp % mod
	default:
		tmp := splitCount(n/2 - 1)
		return tmp * 100 % mod * tmp % mod
	}
}
