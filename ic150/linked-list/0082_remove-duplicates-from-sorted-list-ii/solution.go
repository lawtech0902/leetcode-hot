/*
__author__ = 'robin-luo'
__date__ = '2024/01/02 16:19'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}

		return deleteDuplicates(head.Next)
	}

	head.Next = deleteDuplicates(head.Next)
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	t := dummy
	for t.Next != nil && t.Next.Next != nil {
		if t.Next.Val == t.Next.Next.Val {
			x := t.Next.Val
			for t.Next != nil && t.Next.Val == x {
				t.Next = t.Next.Next
			}
		} else {
			t = t.Next
		}
	}

	return dummy.Next
}
