/*
 * Author: robin-luo
 * Created time: 2024-02-28 15:09:24
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-17 21:02:00
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursion
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

// iteration
func reverseKGroup1(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	cur := head
	size := 0
	for cur != nil {
		size++
		cur = cur.Next
	}

	var pre *ListNode
	cur = head
	for size >= k {
		for i := 0; i < k; i++ {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}

		next := p.Next
		p.Next.Next = cur
		p.Next = pre
		p = next
		size -= k
	}

	return dummy.Next
}
