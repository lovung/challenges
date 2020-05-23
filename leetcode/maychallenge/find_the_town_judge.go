package maychallenge

func findJudge(N int, trust [][]int) int {
	except := make([]int, N)
	count := make([]int, N)
	length := len(trust)
	for i := 0; i < length; i++ {
		count[trust[i][1]-1]++
		except[trust[i][0]-1] = 1
	}
	for i := 0; i < N; i++ {
		if except[i] == 0 && count[i] == N-1 {
			return i + 1
		}
	}
	return -1
}