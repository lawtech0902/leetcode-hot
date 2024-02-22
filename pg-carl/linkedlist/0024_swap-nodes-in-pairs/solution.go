/*
 * Author: robin-luo
 * Created time: 2024-02-21 16:16:57
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 16:33:11
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
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		temp1 := cur.Next
		temp2 := cur.Next.Next
		cur.Next = temp2
		temp1.Next = temp2.Next
		temp2.Next = temp1
		cur = temp1
	}

	return dummy.Next
}
