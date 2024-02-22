/*
 * Author: robin-luo
 * Created time: 2024-02-21 10:36:31
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 11:49:06
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var (
		s1, s2           []int
		a, b, cur, carry int
		res              *ListNode
	)

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		if len(s1) == 0 {
			a = 0
		} else {
			a = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) == 0 {
			b = 0
		} else {
			b = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		cur = a + b + carry
		cur, carry = cur%10, cur/10
		curNode := &ListNode{Val: cur}
		curNode.Next = res
		res = curNode
	}

	return res
}
