/*
__author__ = 'robin-luo'
__date__ = '2025/06/23 19:05'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}

			return head
		}
	}

	return nil
}
