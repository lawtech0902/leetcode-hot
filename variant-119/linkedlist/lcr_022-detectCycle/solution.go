/*
 * Author: robin-luo
 * Created time: 2024-02-21 10:18:09
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 10:20:11
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
			for slow != dummy {
				slow = slow.Next
				dummy = dummy.Next
			}

			return dummy
		}
	}

	return nil
}
