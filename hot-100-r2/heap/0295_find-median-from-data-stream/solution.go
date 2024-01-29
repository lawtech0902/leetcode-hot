/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 11:27'
*/

package solution

import (
	"container/heap"
)

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

func Constructor() MedianFinder {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() != this.maxHeap.Len() {
		heap.Push(this.minHeap, num)
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	} else {
		heap.Push(this.maxHeap, num)
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() != this.maxHeap.Len() {
		return float64((*this.minHeap)[0]+(*this.maxHeap)[0]) / 2.0
	} else {
		return float64((*this.minHeap)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

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

type MaxHeap []int

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MaxHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *MaxHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}
