/*
 * Author: robin-luo
 * Created time: 2024-02-21 13:45:01
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 13:50:26
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var pre *ListNode
	for slow != nil {
		temp := slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}

	for pre != nil && head != nil {
		if pre.Val != head.Val {
			return false
		}

		pre = pre.Next
		head = head.Next
	}

	return true
}
