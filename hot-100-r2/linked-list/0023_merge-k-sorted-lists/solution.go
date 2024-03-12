/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 13:59'
*/

package solution

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	if m[i] == nil {
		return false
	}

	if m[j] == nil {
		return true
	}

	return m[i].Val < m[j].Val
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

func (m *MinHeap) Pop() any {
	size := len(*m)
	res := (*m)[size-1]
	*m = (*m)[:size-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	mh := &MinHeap{}
	heap.Init(mh)
	for _, l := range lists {
		if l != nil {
			heap.Push(mh, l)
		}
	}

	dummy := &ListNode{}
	p := dummy
	for mh.Len() > 0 {
		node := heap.Pop(mh).(*ListNode)
		if node.Next != nil {
			heap.Push(mh, node.Next)
		}

		p.Next = node
		p = p.Next
	}

	return dummy.Next
}
