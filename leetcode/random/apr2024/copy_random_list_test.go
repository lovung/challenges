package apr2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_copyRandomList(t *testing.T) {
	node4 := Node{Val: 1}
	node3 := Node{Val: 10, Next: &node4}
	node2 := Node{Val: 11, Next: &node3, Random: &node4}
	node1 := Node{Val: 13, Next: &node2}
	head := Node{Val: 7, Next: &node1}
	node1.Random = &head
	node3.Random = &node2
	node4.Random = &head
	got := copyRandomList(&head)
	m := got
	for n := &head; n != nil; n = n.Next {
		assert.Equal(t, n.Val, m.Val)
		if n.Random == nil {
			assert.Nil(t, m.Random)
		} else {
			assert.Equal(t, n.Random.Val, m.Random.Val)
		}
		m = m.Next
	}
}
