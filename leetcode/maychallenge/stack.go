package maychallenge

import "fmt"
import "errors"

// Stack the stack to save k first node
type Stack struct {
	size  int
	count int
	stack []interface{}
}

// NewStack create new stack with
func NewStack(n int) *Stack {
	return &Stack{
		size:  n,
		count: 0,
		stack: make([]interface{}, n),
	}
}

// func (s *Stack) print() {
// 	for i := 0; i < s.count; i++ {
// 		fmt.Printf("%v->", s.stack[i])
// 	}
// 	fmt.Println()
// }

func (s *Stack) push(n interface{}) error {
	if s.count == s.size {
		return errors.New("Full")
	}
	s.stack[s.count] = n
	s.count++
	return nil
}

func (s *Stack) pop() interface{} {
	if s.count == 0 {
		return nil
	}
	out := s.stack[s.count-1]
	s.stack[s.count-1] = nil
	s.count--
	return out
}

// func (s *Stack) top() interface{} {
// 	if s.count == 0 {
// 		return nil
// 	}
// 	return s.stack[s.count-1]
// }

// func (s *Stack) bottom() interface{} {
// 	if len(s.stack) == 0 {
// 		return nil
// 	}
// 	return s.stack[0]
// }

// func (s *Stack) getCount() int {
// 	return s.count
// }

// func (s *Stack) getSize() int {
// 	return s.size
// }
