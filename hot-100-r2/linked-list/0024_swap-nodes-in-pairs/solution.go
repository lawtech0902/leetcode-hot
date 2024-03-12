/*
__author__ = 'robin-luo'
__date__ = '2024/01/23 16:46'
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
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := pre.Next.Next
		pre.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		pre = node1
	}

	return dummy.Next
}
