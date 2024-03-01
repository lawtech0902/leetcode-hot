/*
 * Author: robin-luo
 * Created time: 2024-02-29 15:39:41
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 15:40:59
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
