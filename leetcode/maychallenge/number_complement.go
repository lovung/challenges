package maychallenge

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3319/
 */

func findComplement(num int) int {
	// Solution 1
	// tmp := num
	// counter := 1
	// var output int
	// for tmp > 0 {
	//     if tmp & 0x1 == 0 {
	//         output += counter
	//     }
	//     counter <<= 1
	//     tmp >>= 1
	// }
	// return output

	// Solution 2
	x := ^0
	for ; x&num > 0; x <<= 1 {
	}
	return ^x ^ num
}
