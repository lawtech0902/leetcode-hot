/*
__author__ = 'robin-luo'
__date__ = '2023/11/01 10:51'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	prev.Next = prev.Next.Next
	return head
}
