/*
 * Author: robin-luo
 * Created time: 2024-03-06 10:38:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-06 10:43:43
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

	p := head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}
