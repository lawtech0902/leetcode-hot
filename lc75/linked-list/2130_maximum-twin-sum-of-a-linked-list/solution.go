/*
__author__ = 'robin-luo'
__date__ = '2023/11/01 11:32'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head.Next
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	last := slow.Next
	for last.Next != nil {
		curr := last.Next
		last.Next = curr.Next
		curr.Next = slow.Next
		slow.Next = curr
	}

	res := 0
	x, y := head, slow.Next
	for y != nil {
		res = max(res, x.Val+y.Val)
		x = x.Next
		y = y.Next
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
