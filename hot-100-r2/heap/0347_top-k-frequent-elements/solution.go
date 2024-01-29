/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 10:54'
*/

package solution

import "container/heap"

type Element struct {
	num, count int
}

type PriorityQueue []*Element

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Element))
}

func (pq *PriorityQueue) Pop() any {
	element := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return element
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	pq := make(PriorityQueue, 0)
	for num, count := range countMap {
		pq = append(pq, &Element{
			num:   num,
			count: count,
		})
	}

	heap.Init(&pq)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&pq).(*Element).num
	}

	return res
}
