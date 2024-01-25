/*
__author__ = 'robin-luo'
__date__ = '2024/01/23 15:26'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	i := 1
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}

		i++
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
