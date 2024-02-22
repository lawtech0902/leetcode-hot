/*
 * Author: robin-luo
 * Created time: 2024-02-21 15:56:13
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 15:58:31
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{-1, head}
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
