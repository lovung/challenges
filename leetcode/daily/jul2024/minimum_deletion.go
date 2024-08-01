package jul2024

// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/
func minimumDeletions(s string) int {
	cntB, cntA := 0, 0
	bs := []byte(s)
	storeCntA := make([]int, len(bs))
	storeCntB := make([]int, len(bs))
	for i, b := range bs {
		if b == 'b' {
			cntB++
		}
		storeCntB[i] = cntB
	}
	for i := len(bs) - 1; i >= 0; i-- {
		storeCntA[i] = cntA
		if bs[i] == 'a' {
			cntA++
		}
	}
	res := cntA // maximum we can remove all 'a' char
	for i := range s {
		res = min(res, storeCntB[i]+storeCntA[i])
	}
	return res
}
