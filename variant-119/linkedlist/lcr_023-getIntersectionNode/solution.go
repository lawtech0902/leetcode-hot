/*
 * Author: robin-luo
 * Created time: 2024-02-21 10:22:02
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 10:24:47
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA != nil && headB != nil {
		p1, p2 := headA, headB
		for p1 != p2 {
			if p1 != nil {
				p1 = p1.Next
			} else {
				p1 = headB
			}

			if p2 != nil {
				p2 = p2.Next
			} else {
				p2 = headA
			}
		}

		return p1
	}

	return nil
}
