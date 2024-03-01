/*
 * Author: robin-luo
 * Created time: 2024-03-01 10:32:44
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 10:37:43
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
			dummy := head
			for dummy != slow {
				dummy = dummy.Next
				slow = slow.Next
			}

			return dummy
		}
	}

	return nil
}
