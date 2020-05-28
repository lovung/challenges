package maychallenge

// import "errors"

// func possibleBipartition(N int, dislikes [][]int) bool {
// 	var (
// 		m     = len(dislikes)
// 		gA    = make([]bool, N+1)
// 		gB    = make([]bool, N+1)
// 		queue = Queue{}
// 		first = true
// 	)
// 	for i := 0; i < m; i++ {
// 		err := process(gA, gB, dislikes[i], &queue, &first)
// 		if err != nil {
// 			return false
// 		}
// 	}
// 	for queue.count() > 0 {
// 		pair := queue.peek()
// 		err := process2(gA, gB, pair, &queue)
// 		if err != nil {
// 			return false
// 		}
// 		queue.pop()
// 	}
// 	return true

// }

// func process(gA, gB []bool, pair []int, queue *Queue, first *bool) error {
// 	if gA[pair[0]] && gA[pair[1]] {
// 		return errors.New("error")
// 	}
// 	if gB[pair[0]] && gB[pair[1]] {
// 		return errors.New("error")
// 	}

// 	if !gA[pair[0]] && !gB[pair[0]] && !gA[pair[1]] && !gB[pair[1]] {
// 		if *first {
// 			gA[pair[0]] = true
// 			gB[pair[1]] = true
// 			*first = false
// 		} else {
// 			queue.push(pair)
// 		}
// 		return nil
// 	}
// 	if gA[pair[0]] {
// 		gB[pair[1]] = true
// 	} else if gB[pair[0]] {
// 		gA[pair[1]] = true
// 	} else if gA[pair[1]] {
// 		gB[pair[0]] = true
// 	} else if gB[pair[1]] {
// 		gA[pair[0]] = true
// 	} else {
// 		return errors.New("error")
// 	}
// 	return nil
// }

// func process2(gA, gB []bool, pair []int, queue *Queue) error {
// 	if gA[pair[0]] && gA[pair[1]] {
// 		return errors.New("error")
// 	}
// 	if gB[pair[0]] && gB[pair[1]] {
// 		return errors.New("error")
// 	}

// 	if gA[pair[0]] {
// 		gB[pair[1]] = true
// 	} else if gB[pair[0]] {
// 		gA[pair[1]] = true
// 	} else if gA[pair[1]] {
// 		gB[pair[0]] = true
// 	} else if gB[pair[1]] {
// 		gA[pair[0]] = true
// 	}
// 	return nil
// }

// // Queue ...
// type Queue [][]int

// func (s Queue) peek() []int {
// 	return s[0]
// }

// func (s Queue) count() int {
// 	return len(s)
// }

// func (s *Queue) pop() {
// 	*s = (*s)[1:]
// }

// func (s *Queue) push(value []int) {
// 	*s = append(*s, value)
// }
