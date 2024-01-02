/*
__author__ = 'robin-luo'
__date__ = '2024/01/02 11:42'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	p := dummy
	q := head

	for i := 0; i < m-1; i++ {
		p = p.Next
		q = q.Next
	}

	var (
		tmp, start *ListNode
	)
	end := q

	for i := m; i < n+1; i++ {
		tmp = q.Next
		q.Next = start
		start = q
		q = tmp
	}

	p.Next = start
	end.Next = tmp
	return dummy.Next
}
