package biweekly129

func numberOfStableArrays1(zero int, one int, limit int) int {
	cnt := 0
	recursive(&cnt, zero, one, limit, step{0, 0, 0, 0})
	return cnt
}

func recursive(cnt *int, zero, one, limit int, s step) {
	if zero == s.zeroCnt && one == s.oneCnt {
		*cnt++
		*cnt %= 1e9 + 7
		// fmt.Println(*cnt, s.prev)
		return
	}
	if zero > s.zeroCnt {
		if s.last != 0 || s.similarCnt < limit {
			similarCnt := s.similarCnt
			if s.last != 0 {
				similarCnt = 0
			}
			recursive(cnt, zero, one, limit, step{
				s.zeroCnt + 1, s.oneCnt, similarCnt + 1, 0, // append(s.prev, 0),
			})
		}
	}
	if one > s.oneCnt {
		if s.last != 1 || s.similarCnt < limit {
			similarCnt := s.similarCnt
			if s.last != 1 {
				similarCnt = 0
			}
			recursive(cnt, zero, one, limit, step{
				s.zeroCnt, s.oneCnt + 1, similarCnt + 1, 1, // append(s.prev, 1),
			})
		}
	}
}

type step struct {
	zeroCnt    int
	oneCnt     int
	similarCnt int
	last       int
	// prev       []int
}
