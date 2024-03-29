/*
 * Author: robin-luo
 * Created time: 2024-02-28 10:14:49
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-15 11:04:05
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortOddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := partition(head)
	even = reverse(even)
	return merge(odd, even)
}

func partition(head *ListNode) (*ListNode, *ListNode) {
	evenHead := head.Next
	odd, even := head, evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	return head, evenHead
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}

		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return dummy.Next
}
