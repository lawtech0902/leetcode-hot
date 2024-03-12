/*
__author__ = 'robin-luo'
__date__ = '2024/01/22 16:29'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := head.Next
	newHead := reverseList(tail)
	head.Next = nil
	tail.Next = head
	return newHead
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}
