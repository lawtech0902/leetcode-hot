/*
 * Author: robin-luo
 * Created time: 2024-02-28 15:41:29
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 17:30:56
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Next: head}
	revPre := dummy
	count := 1
	for count < left {
		revPre = revPre.Next
		count++
	}

	revHead := revPre.Next
	var pre *ListNode
	cur := revHead
	for count <= right {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
		count++
	}

	revPre.Next = pre
	revHead.Next = cur
	return dummy.Next
}
