/*
__author__ = 'robin-luo'
__date__ = '2025/06/24 10:36'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		temp1 := p.Next
		temp2 := p.Next.Next
		p.Next = temp2
		temp1.Next = temp2.Next
		temp2.Next = temp1
		p = temp1
	}

	return dummy.Next
}
