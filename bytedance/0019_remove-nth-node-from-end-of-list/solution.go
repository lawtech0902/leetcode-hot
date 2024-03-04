/*
 * Author: robin-luo
 * Created time: 2024-03-01 16:21:58
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 16:43:15
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	slow := dummy
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
