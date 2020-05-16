package maychallenge

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var (
		head = &ListNode{
			Val:  nums[0],
			Next: nil,
		}
		last = head
	)
	for i := 1; i < len(nums); i++ {
		newNode := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		last.Next = newNode
		last = newNode
	}
	return head
}