package jul2024

// https://leetcode.com/problems/find-the-winner-of-the-circular-game/
func findTheWinner(n int, k int) int {
	type Node struct {
		Val  int
		Next *Node
	}
	head := &Node{Val: 1}
	cur := head
	for i := 2; i <= n; i++ {
		newNode := &Node{Val: i}
		cur.Next = newNode
		cur = newNode
	}
	cur.Next = head
	prev := cur
	cur = head
	for range n - 1 {
		for range k - 1 {
			cur, prev = cur.Next, cur
		}
		prev.Next = cur.Next
		cur = cur.Next
	}
	return cur.Val
}
