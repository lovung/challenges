package mathematics

/*
 * Complete the findPoint function below.
 * Link: https://www.hackerrank.com/challenges/find-point/problem
 */
func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
	return []int32{qx*2 - px, qy*2 - py}
}
