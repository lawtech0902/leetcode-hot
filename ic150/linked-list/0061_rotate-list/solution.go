/*
__author__ = 'robin-luo'
__date__ = '2024/01/04 10:49'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	p := dummy
	size := 0
	for p.Next != nil {
		p = p.Next
		size++
	}

	p.Next = dummy.Next
	step := size - (k % size)
	for i := 0; i < step; i++ {
		p = p.Next
	}

	head = p.Next
	p.Next = nil
	return head
}
