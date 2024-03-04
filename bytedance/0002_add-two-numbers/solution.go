/*
 * Author: robin-luo
 * Created time: 2024-02-27 17:44:15
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 15:18:41
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		p.Next = &ListNode{Val: sum}
		p = p.Next
	}

	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
