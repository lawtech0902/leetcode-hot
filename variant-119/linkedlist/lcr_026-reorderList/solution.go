/*
 * Author: robin-luo
 * Created time: 2024-02-21 11:48:52
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 13:44:34
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head1 := slow.Next
	slow.Next = nil
	cur := head1
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	cur1, cur2 := head, pre
	for cur2 != nil {
		temp1, temp2 := cur1.Next, cur2.Next
		cur1.Next = cur2
		cur2.Next = temp1
		cur1, cur2 = temp1, temp2
	}
}
