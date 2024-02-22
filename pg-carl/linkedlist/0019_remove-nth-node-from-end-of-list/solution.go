/*
 * Author: robin-luo
 * Created time: 2024-02-21 16:41:18
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 16:52:25
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	fast := head
	i := 0
	for ; i < n; i++ {
		if fast == nil {
			break
		}

		fast = fast.Next
	}

	if i < n {
		return head
	}

	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
