/*
__author__ = 'robin-luo'
__date__ = '2024/01/23 14:35'
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
	for l1 != nil || l2 != nil || carry > 0 {
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

	return dummy.Next
}
