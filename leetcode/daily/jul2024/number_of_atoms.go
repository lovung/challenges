package jul2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/v2/stacks"
	"github.com/emirpasic/gods/v2/stacks/linkedliststack"
)

// https://leetcode.com/problems/number-of-atoms/
func countOfAtoms(formula string) string {
	sm := NewStateMachine(formula)
	for i := 0; i < sm.Len(); i++ {
		sm.Process(&i)
	}
	return sm.PeakGroup().String()
}

type State int

const (
	Open State = iota
	Close
)

type StateMachine struct {
	formula string
	curEle  string
	state   State
	stack   stacks.Stack[*ElementGroup]
}

func NewStateMachine(formula string) *StateMachine {
	s := linkedliststack.New[*ElementGroup]()
	globalMap := make(ElementGroup)
	s.Push(&globalMap)
	return &StateMachine{
		formula: formula + "*",
		stack:   s,
	}
}

func (sm *StateMachine) Len() int {
	return len(sm.formula)
}

func (sm *StateMachine) PeakGroup() *ElementGroup {
	val, _ := sm.stack.Peek()
	return val
}

func (sm *StateMachine) Process(i *int) {
	switch sm.state {
	case Open:
		sm.processOpen(i)
	case Close:
		sm.processClose(i)
	}
}

func (sm *StateMachine) processOpen(i *int) {
	r := sm.formula[*i]
	curGroup, _ := sm.stack.Peek()
	switch {
	case r == '(':
		curGroup.Incr(sm.curEle, 1)
		sm.curEle = ""
		newGroup := make(ElementGroup)
		sm.stack.Push(&newGroup)
	case r == ')':
		curGroup.Incr(sm.curEle, 1)
		sm.curEle = ""
		sm.state = Close
	case '0' <= r && r <= '9':
		num := sm.scanNumber(i)
		curGroup.Incr(sm.curEle, num)
		sm.curEle = ""
	default:
		curGroup.Incr(sm.curEle, 1)
		sm.curEle = sm.scanElement(i)
	}
}

func (sm *StateMachine) processClose(i *int) {
	r := sm.formula[*i]
	curGroup, _ := sm.stack.Pop()
	parentGroup, _ := sm.stack.Peek()
	sm.state = Open
	switch {
	case r == '(':
		parentGroup.MergeWithK(*curGroup, 1)
		newGroup := make(ElementGroup)
		sm.stack.Push(&newGroup)
	case r == ')':
		parentGroup.MergeWithK(*curGroup, 1)
		sm.state = Close
	case '0' <= r && r <= '9':
		num := sm.scanNumber(i)
		parentGroup.MergeWithK(*curGroup, num)
	default:
		parentGroup.MergeWithK(*curGroup, 1)
		sm.curEle = sm.scanElement(i)
	}
}

func (sm *StateMachine) scanElement(i *int) string {
	j := *i + 1
	for j < len(sm.formula) && 'a' <= sm.formula[j] && sm.formula[j] <= 'z' {
		j++
	}
	ele := sm.formula[*i:j]
	*i = j - 1
	return ele
}

func (sm *StateMachine) scanNumber(i *int) int {
	j := *i + 1
	for j < len(sm.formula) && '0' <= sm.formula[j] && sm.formula[j] <= '9' {
		j++
	}
	num, _ := strconv.Atoi(sm.formula[*i:j])
	*i = j - 1
	return num
}

type ElementGroup map[string]int

func (m ElementGroup) Incr(e string, cnt int) {
	if len(e) > 0 {
		m[e] += cnt
	}
}

func (m ElementGroup) MergeWithK(n ElementGroup, k int) {
	for key, v := range n {
		m[key] += v * k
	}
}

func (m ElementGroup) String() string {
	type element struct {
		name string
		cnt  int
	}
	elements := make([]*element, 0, len(m))
	for k, v := range m {
		elements = append(elements, &element{k, v})
	}
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].name < elements[j].name
	})
	sb := strings.Builder{}
	for _, e := range elements {
		sb.WriteString(e.name)
		if e.cnt > 1 {
			sb.WriteString(fmt.Sprint(e.cnt))
		}
	}
	return sb.String()
}
