/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 10:28'
*/

package solution

import "container/heap"

type MinHeap []int

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *MinHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	mh := &MinHeap{}
	heap.Init(mh)
	for i := 0; i < k; i++ {
		heap.Push(mh, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < (*mh)[0] {
			continue
		}

		heap.Pop(mh)
		heap.Push(mh, nums[i])
	}

	return (*mh)[0]
}
