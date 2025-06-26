/*
__author__ = 'robin-luo'
__date__ = '2025/06/26 09:16'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := head
	for i := 0; i < k; i++ {
		if dummy == nil {
			return head
		}

		dummy = dummy.Next
	}

	var pre *ListNode
	cur := head
	for i := 0; i < k; i++ {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	head.Next = reverseKGroup(cur, k)
	return pre
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, dummy
	size := 0
	for cur.Next != nil {
		size++
		cur = cur.Next
	}

	for size >= k {
		cur = pre.Next
		for i := 1; i < k; i++ {
			temp := cur.Next
			cur.Next = temp.Next
			temp.Next = pre.Next
			pre.Next = temp
		}

		pre = cur
		size -= k
	}

	return dummy.Next
}
