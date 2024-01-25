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
	temp := MinHeap(lists)
	h := &temp
	heap.Init(h)
	var head, tail *ListNode
	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		if v == nil {
			continue
		}

		if tail != nil {
			tail.Next = v
		}

		if head == nil {
			head = v
		}

		tail = v
		if v.Next != nil {
			heap.Push(h, v.Next)
		}
	}

	return head
}
