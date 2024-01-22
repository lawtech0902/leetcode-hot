/*
__author__ = 'robin-luo'
__date__ = '2024/01/22 17:42'
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
