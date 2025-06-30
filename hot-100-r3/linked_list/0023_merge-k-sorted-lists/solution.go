/*
__author__ = 'robin-luo'
__date__ = '2025/06/30 10:52'
*/

package solution

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i] == nil {
		return false
	}

	if h[j] == nil {
		return true
	}

	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	size := len(*h)
	res := (*h)[size-1]
	*h = (*h)[:size-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	temp := MinHeap(lists)
	h := &temp
	heap.Init(h)
	var head, last *ListNode
	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		if v == nil {
			continue
		}

		if last != nil {
			last.Next = v
		}

		if head == nil {
			head = v
		}

		last = v
		if v.Next != nil {
			heap.Push(h, v.Next)
		}
	}

	return head
}
