/*
 * Author: robin-luo
 * Created time: 2024-02-28 15:29:47
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 15:35:36
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
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
