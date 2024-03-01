/*
 * Author: robin-luo
 * Created time: 2024-02-27 17:47:36
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 14:54:37
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	h1, h2 := head, slow.Next
	slow.Next = nil

	cur := h2
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	cur1, cur2 := h1, pre
	for cur2 != nil {
		temp1, temp2 := cur1.Next, cur2.Next
		cur1.Next = cur2
		cur2.Next = temp1
		cur1, cur2 = temp1, temp2
	}
}
