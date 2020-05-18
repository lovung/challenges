package mathematics

/*
 * Complete the summingSeries function below.
 * https://www.hackerrank.com/challenges/summing-the-n-series/problem
 */
func summingSeries(n int64) int32 {
	tmp := n % 1000000007
	return int32((tmp * tmp) % 1000000007)
}
