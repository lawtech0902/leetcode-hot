/*
 * Author: robin-luo
 * Created time: 2024-02-28 09:55:56
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 10:02:46
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
