/*
 * Author: robin-luo
 * Created time: 2024-03-19 11:23:26
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 14:59:41
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	cur := head
	for cur != nil {
		if cur.Next.Val < cur.Val {
			pre := dummy
			// 寻找插入点
			for pre.Next.Val < cur.Next.Val {
				pre = pre.Next
			}

			temp := cur.Next
			cur.Next = temp.Next
			temp.Next = pre.Next
			pre.Next = temp
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
