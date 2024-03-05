/*
 * Author: robin-luo
 * Created time: 2024-03-04 15:55:44
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 16:01:13
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	slow, fast := head, head
	for i := 0; i < cnt; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
