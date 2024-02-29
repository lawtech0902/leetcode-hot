/*
 * Author: robin-luo
 * Created time: 2024-02-28 13:57:05
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 14:12:34
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
	dummy := &ListNode{Next: head}
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			v := p.Next.Val
			for p.Next != nil && p.Next.Val == v {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}

	return dummy.Next
}
