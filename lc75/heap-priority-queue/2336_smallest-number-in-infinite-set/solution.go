/*
__author__ = 'robin-luo'
__date__ = '2023/11/23 10:31'
*/

package solution

import (
	"container/heap"
)

type PriorityQueue []int

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type SmallestInfiniteSet struct {
	queue PriorityQueue
	index int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		queue: make(PriorityQueue, 0),
		index: 0,
	}
}

func (sis *SmallestInfiniteSet) PopSmallest() int {
	if len(sis.queue) != 0 {
		if sis.queue[0] > sis.index {
			sis.index++
			return sis.index
		}
		return heap.Pop(&sis.queue).(int)
	}
	sis.index++
	return sis.index
}

func (sis *SmallestInfiniteSet) AddBack(num int) {
	if !sis.contains(num) && sis.index >= num {
		heap.Push(&sis.queue, num)
	}
}

func (sis *SmallestInfiniteSet) contains(num int) bool {
	for _, n := range sis.queue {
		if n == num {
			return true
		}
	}
	return false
}
