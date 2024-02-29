/*
 * Author: robin-luo
 * Created time: 2024-02-28 13:35:09
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 14:12:34
 */

package solution

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Val < mh[j].Val
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(*ListNode))
}

func (mh *MinHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	mh := MinHeap{}
	for _, l := range lists {
		if l != nil {
			mh = append(mh, l)
		}
	}

	heap.Init(&mh)
	dummy := &ListNode{}
	p := dummy
	for mh.Len() > 0 {
		node := heap.Pop(&mh).(*ListNode)
		if node.Next != nil {
			heap.Push(&mh, node.Next)
		}

		p.Next = node
		p = p.Next
	}

	return dummy.Next
}
